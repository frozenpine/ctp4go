package state

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"
)

const (
	DefaultBackoffDuration = time.Second * 3
	DefaultMaxBackOff      = time.Minute
)

type HandlerOpt func(*Handler)

type Handler struct {
	deadline       time.Duration
	retry          int
	timeout        time.Duration
	recoverTimeout bool
	backoff        func(time.Duration, time.Duration, int) time.Duration

	hdl       func() error
	execCount int

	next *Handler
}

func NewHandler(fn func() error, options ...HandlerOpt) (*Handler, error) {
	if fn == nil {
		return nil, errors.New("handler is nil")
	}

	hdl := Handler{
		hdl: fn,
		backoff: func(d1, d2 time.Duration, i int) time.Duration {
			dur := min(d1*time.Duration(i), d2)

			return dur + time.Duration(rand.Int63n(int64(d1/10)))
		},
	}

	return &hdl, nil
}

func (h *Handler) SetNext(next *Handler) *Handler {
	if h.next != nil {
		return h.next.SetNext(next)
	}

	h.next = next
	return h.next
}

func (h *Handler) Handle() error {
	deadlineCtx := context.Background()

	if h.deadline > 0 {
		var deadlineCancel context.CancelFunc

		deadlineCtx, deadlineCancel = context.WithDeadline(
			deadlineCtx, time.Now().Add(h.deadline),
		)
		defer deadlineCancel()
	}

TRY:
	for range h.retry + 1 {
		timeoutCtx := deadlineCtx
		if h.timeout > 0 {
			var timeoutCancel context.CancelFunc

			timeoutCtx, timeoutCancel = context.WithTimeout(
				timeoutCtx, h.timeout,
			)

			defer timeoutCancel()
		}

		waitHandle := make(chan error)

		go func() {
			defer func() {
				if err := recover(); err != nil {
					waitHandle <- fmt.Errorf(
						"handler run failed: %+v\n%s",
						err, string(debug.Stack()),
					)
				}
			}()
			waitHandle <- h.hdl()
		}()

		select {
		case <-deadlineCtx.Done():
			return deadlineCtx.Err()
		case <-timeoutCtx.Done():
			err := timeoutCtx.Err()

			if !h.recoverTimeout {
				return err
			}
		case err := <-waitHandle:
			h.execCount++

			if err == nil {
				break TRY
			}
		}

		<-time.After(h.backoff(
			DefaultBackoffDuration,
			DefaultMaxBackOff,
			h.execCount,
		))
	}

	if h.next != nil {
		return h.next.Handle()
	}

	return nil
}
