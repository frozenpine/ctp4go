package state

import (
	"context"
	"fmt"
	"log/slog"
	"reflect"
	"sync"
	"time"
	"weak"

	"github.com/frozenpine/ctp4go"
)

type Flag[T comparable, V any] interface {
	Name() string
	SetFlag(T, ...V) error
	GetFlag() T
	GetFlagWithPayload() (T, V)
	GetPayload() V
	SetPayload(V)
	CheckFlag(v T) bool
	Wait(T, time.Duration) error

	RLock()
	RUnlock()
	Lock()
	Unlock()
	fmt.Stringer
}

type notify struct {
	sync.Once
	ch chan struct{}
}

func (n *notify) close() {
	n.Do(func() {
		close(n.ch)
	})
}

type BaseFlag[T comparable, V any] struct {
	sync.RWMutex

	name     string
	flag     T
	migrator func(*T, T) error
	payload  V

	notifies map[weak.Pointer[notify]]struct{}
}

type flagInitOpt[T comparable, V any] func(*BaseFlag[T, V])

func WithInit[T comparable, V any](v T) flagInitOpt[T, V] {
	return func(f *BaseFlag[T, V]) {
		f.flag = v
	}
}

// NewBaseFlag 创建命名旗标，可为任意可比较类型
//
// 若类型实现了 `Migrate(T) error` 接口，则设置旗标值时将使用该接口
func NewBaseFlag[T comparable, V any](
	name string, options ...flagInitOpt[T, V],
) (flag *BaseFlag[T, V]) {
	flagType := reflect.TypeFor[T]()

	if flagType.Kind() != reflect.Pointer {
		flagType = reflect.PointerTo(flagType)
	}

	flag = &BaseFlag[T, V]{
		name:     name,
		notifies: map[weak.Pointer[notify]]struct{}{},
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		opt(flag)
	}

	if fn, ok := flagType.MethodByName("Migrate"); ok {
		flag.migrator = fn.Func.Interface().(func(*T, T) error)
	}

	return
}

func (s *BaseFlag[T, V]) makeNoitfy() (<-chan struct{}, func()) {
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

func (s *BaseFlag[T, V]) Name() string { return s.name }

func (s *BaseFlag[T, V]) String() string {
	buff := ctp4go.GetStringBuilder()
	fmt.Fprintf(buff, "%s.%+v", s.Name(), s.flag)
	return buff.String()
}

func (s *BaseFlag[T, V]) notifyAll() {
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
}

// SetFlag 设置当前标记值
func (s *BaseFlag[T, V]) SetFlag(v T, payload ...V) error {
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

	if len(payload) > 0 {
		s.payload = payload[0]
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"state flag setted, find waiting",
		slog.String("name", s.name),
		slog.Any("flag", v),
	)

	s.notifyAll()
	return nil
}

// GetFlag 获取当前标记值
func (s *BaseFlag[T, V]) GetFlag() T {
	s.RLock()
	defer s.RUnlock()

	return s.flag
}

// GetFlag 获取当前标记值及状态负载
func (s *BaseFlag[T, V]) GetFlagWithPayload() (T, V) {
	s.RLock()
	defer s.RUnlock()

	return s.flag, s.GetPayload()
}

// GetPayload 获取当前状态负载
func (s *BaseFlag[T, V]) GetPayload() V {
	s.RLock()
	defer s.RUnlock()

	return s.payload
}

func (s *BaseFlag[T, V]) SetPayload(v V) {
	s.Lock()
	defer s.Unlock()

	s.payload = v
}

// CheckFlag 检测当前标记值是否为指定值
func (s *BaseFlag[T, V]) CheckFlag(v T) bool {
	s.RLock()
	defer s.RUnlock()

	return s.flag == v
}

// Wait 同步等待指定标记值
// 超时最小间隔为1s，小于等于0则无超时
func (s *BaseFlag[T, V]) Wait(v T, timeout time.Duration) error {
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

	for !s.CheckFlag(v) {
		select {
		case <-tm.Done():
			if s.CheckFlag(v) {
				return nil
			}

			return tm.Err()
		case <-time.Tick(time.Second):
			if s.CheckFlag(v) {
				return nil
			}
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
	}

	return nil
}
