package trader

import (
	"context"
	"sync"
	"time"
)

type State[T comparable] struct {
	sync.RWMutex

	flag T
}

func (s *State[T]) SetFlag(v T) {
	s.Lock()
	defer s.Unlock()

	s.flag = v
}

func (s *State[T]) GetFlag() T {
	s.RLock()
	defer s.RLock()

	return s.flag
}

func (s *State[T]) Wait(v T, timeout time.Duration) error {
	wait := context.Background()
	if timeout = timeout.Round(time.Second); timeout > 0 {
		var cancel context.CancelFunc

		wait, cancel = context.WithTimeout(wait, timeout)

		defer cancel()
	}

	for {
		select {
		case <-wait.Done():
			if s.GetFlag() == v {
				return nil
			}

			return wait.Err()
		case <-time.Tick(time.Second):
			if s.GetFlag() == v {
				return nil
			}
		}
	}
}
