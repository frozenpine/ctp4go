package cache

import (
	"context"
	"log/slog"
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

type State[T comparable] struct {
	sync.RWMutex

	name     string
	flag     T
	notifies map[weak.Pointer[notify]]struct{}
}

func NewState[T comparable](name string) *State[T] {
	return &State[T]{
		name:     name,
		notifies: map[weak.Pointer[notify]]struct{}{},
	}
}

func (s *State[T]) makeNoitfy() (<-chan struct{}, func()) {
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

func (s *State[T]) SetFlag(v T) {
	s.Lock()
	defer s.Unlock()

	s.flag = v

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
}

func (s *State[T]) GetFlag() T {
	s.RLock()
	defer s.RUnlock()

	return s.flag
}

func (s *State[T]) Wait(v T, timeout time.Duration) error {
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
