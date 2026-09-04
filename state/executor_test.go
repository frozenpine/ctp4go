package state_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/state"
)

func TestExecutor(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 1)

	chain, err := state.NewExecutor(t.Context(), "test")
	if err != nil {
		t.Fatal(err)
	}

	if err := chain.ConcurrentExecute(); err == nil {
		t.Fatal("empty execution assert failed")
	} else {
		t.Log(err)
	}

	if err := chain.ConcurrentExecute(
		state.WithExecSpan(func(s *state.SpanInfo) error {
			go func() {
				<-time.After(time.Second)
				s.SetFlag(state.SpanRspFinished)
				t.Log("first rsponse done")
			}()
			t.Log("first executed")
			return nil
		}),
		state.WithExecSpan(func(s *state.SpanInfo) error {
			go func() {
				<-time.After(time.Second)
				s.SetFlag(state.SpanRspFinished)
				t.Log("last response done")
			}()
			t.Log("last executed")
			return nil
		}),
	); err != nil {
		t.Fatal(err)
	}

	if err := chain.WaitAll(0); err != nil {
		t.Fatal(err)
	}

	// if err := chain.ConcurentExecuteAndWait(
	// 	time.Second*11,
	// 	state.WithExecSpan(func(s *state.SpanInfo) error {
	// 		go func() {
	// 			<-time.After(time.Second)
	// 			s.SetFlag(state.SpanRspFinished)
	// 			t.Log("first rsponse done")
	// 		}()
	// 		t.Log("first executed")
	// 		return nil
	// 	}),
	// 	state.WithExecSpan(func(s *state.SpanInfo) error {
	// 		go func() {
	// 			<-time.After(time.Second)
	// 			s.SetFlag(state.SpanRspFinished)
	// 			t.Log("last response done")
	// 		}()
	// 		t.Log("last executed")
	// 		return nil
	// 	}),
	// ); err != nil {
	// 	t.Fatal(err)
	// } else {
	// 	t.Log("all span executed")
	// }

	// if err := chain.Reset(
	// 	state.WithExecSpan(func(s *state.SpanInfo) error {
	// 		go func() {
	// 			<-time.After(time.Second)
	// 			s.SetFlag(state.SpanRspFinished)
	// 			t.Log("first response done")
	// 		}()
	// 		t.Log("first executed")
	// 		return nil
	// 	}, func(*state.SpanInfo) error {
	// 		<-time.After(time.Second)
	// 		t.Log("first prepare")
	// 		return nil
	// 	}),
	// 	state.WithExecSpan(func(s *state.SpanInfo) error {
	// 		go func() {
	// 			<-time.After(time.Second)
	// 			s.SetFlag(state.SpanBreaked)
	// 			t.Log("last response done")
	// 		}()
	// 		t.Log("last executed")
	// 		return nil
	// 	}),
	// ); err != nil {
	// 	t.Fatal(err)
	// }

	// if err := chain.ChainExecuteAndWait(time.Second * 16); err != nil {
	// 	if !errors.Is(err, state.ErrReqSpanBreaked) {
	// 		t.Fatal(err)
	// 	} else {
	// 		t.Log(err)
	// 	}
	// }
}

func TestSpanInject(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 1)

	chain, err := state.NewExecutor(t.Context(), "test")
	if err != nil {
		t.Fatal(err)
	}

	if err := chain.ChainExecuteAndWait(
		time.Second*20,
		state.WithExecSpan(func(si *state.SpanInfo) error {
			go func() {
				<-time.After(time.Second)
				si.SetFlag(state.SpanRspFinished)
			}()

			t.Log("first executed", time.Now())

			si.InjectSpan(func(si *state.SpanInfo) error {
				go func() {
					<-time.After(2 * time.Second)
					si.SetFlag(state.SpanRspFinished)
				}()
				t.Log("middle executed", time.Now())
				return nil
			})
			return nil
		}),
		state.WithExecSpan(func(si *state.SpanInfo) error {
			go func() {
				<-time.After(3 * time.Second)
				si.SetFlag(state.SpanRspFinished)
			}()
			t.Log("last executed", time.Now())
			return nil
		}),
	); err != nil {
		t.Fatal(err)
	}
}
