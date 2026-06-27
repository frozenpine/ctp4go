package state

import (
	"context"
	"errors"
	"log/slog"
	"reflect"
	"sync"
	"time"
	"weak"
)

type notify struct {
	sync.Once
	ch chan struct{}
}

func (n *notify) close() {
	n.Do(func() {
		close(n.ch)
	})
}

type Flag[T comparable] struct {
	sync.RWMutex

	name     string
	flag     T
	migrator func(*T, T) error

	notifies map[weak.Pointer[notify]]struct{}
}

func NewFlag[T comparable](name string) (flag *Flag[T]) {
	flagType := reflect.TypeFor[T]()

	if flagType.Kind() != reflect.Pointer {
		flagType = reflect.PointerTo(flagType)
	}

	flag = &Flag[T]{
		name:     name,
		notifies: map[weak.Pointer[notify]]struct{}{},
	}

	if fn, ok := flagType.MethodByName(
		"Migrate",
	); ok {
		reflect.ValueOf(&flag.migrator).Elem().Set(reflect.MakeFunc(
			reflect.TypeOf(flag.migrator),
			func(args []reflect.Value) []reflect.Value {
				return fn.Func.Call(args)
			},
		))
	}

	return
}

func (s *Flag[T]) makeNoitfy() (<-chan struct{}, func()) {
	n := &notify{ch: make(chan struct{})}
	wait := weak.Make(n)

	s.Lock()
	defer s.Unlock()

	s.notifies[wait] = struct{}{}

	return n.ch, func() {
		s.Lock()
		defer s.Unlock()

		n.close()
		delete(s.notifies, wait)
	}
}

// SetFlag 设置当前标记值
func (s *Flag[T]) SetFlag(v T) error {
	s.Lock()
	defer s.Unlock()

	if s.migrator != nil {
		err := s.migrator(&s.flag, v)
		if err != nil {
			return err
		}
	} else {
		s.flag = v
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"state flag setted, find waiting",
		slog.String("name", s.name),
		slog.Any("flag", v),
	)

	count := len(s.notifies)
	for w := range s.notifies {
		if n := w.Value(); n != nil {
			slog.Log(
				context.Background(), slog.LevelDebug-2,
				"notify for waiting",
				slog.Any("wait", w),
			)

			select {
			case n.ch <- struct{}{}:
			default:
			}
		} else {
			slog.Log(
				context.Background(), slog.LevelDebug-2,
				"notify waiting recycled",
				slog.Any("wait", n),
			)
		}
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"all waiting notified",
		slog.String("name", s.name),
		slog.Int("count", count),
	)
	return nil
}

// GetFlag 获取当前标记值
func (s *Flag[T]) GetFlag() T {
	s.RLock()
	defer s.RUnlock()

	return s.flag
}

// CheckFlag 检测当前标记值是否为指定值
func (s *Flag[T]) CheckFlag(v T) bool {
	s.RLock()
	defer s.RUnlock()

	return s.flag == v
}

// Wait 同步等待指定标记值
// 超时最小间隔为1s，小于等于0则无超时
func (s *Flag[T]) Wait(v T, timeout time.Duration) error {
	tm := context.Background()
	if timeout = timeout.Round(time.Second); timeout > 0 {
		var tmCancel context.CancelFunc

		tm, tmCancel = context.WithTimeout(tm, timeout)

		defer tmCancel()
	}

	wait, cancel := s.makeNoitfy()
	defer cancel()

	slog.Log(
		tm, slog.LevelDebug-2,
		"start waiting state",
		slog.String("name", s.name),
		slog.Any("cond", v),
	)

	for {
		select {
		case <-tm.Done():
			if s.GetFlag() == v {
				return nil
			}

			return tm.Err()
		case <-time.Tick(time.Second):
			slog.Log(
				tm, slog.LevelDebug-2,
				"wait second tick triggered",
			)
		case <-wait:
			slog.Log(
				tm, slog.LevelDebug-2,
				"notify wait triggered",
			)
		}

		if s.GetFlag() == v {
			return nil
		}
	}
}

type FlagResponsor[T comparable] struct {
	*Flag[T]

	responsors map[T]*Handler
}

func NewFlagResponsor[T comparable](name string) *FlagResponsor[T] {
	return &FlagResponsor[T]{
		Flag:       NewFlag[T](name),
		responsors: make(map[T]*Handler),
	}
}

func (r *FlagResponsor[T]) SetFlag(v T) (err error) {
	if err = r.Flag.SetFlag(v); err == nil {
		r.RLock()
		defer r.RUnlock()

		if rsp, exist := r.responsors[v]; !exist {
			return
		} else {
			return rsp.Handle()
		}
	}

	return
}

func (r *FlagResponsor[T]) AddHandler(v T, hdl *Handler) error {
	if hdl == nil {
		return errors.New("invalid responsor handler")
	}

	r.Lock()
	defer r.Unlock()

	if rsp, exist := r.responsors[v]; !exist {
		r.responsors[v] = hdl
	} else {
		rsp.SetNext(hdl)
	}

	return nil
}
