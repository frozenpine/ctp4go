package trader_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/trader"
)

func TestState(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	state := trader.NewState[bool]("test")

	for idx := range 10 {
		go func() {
			err := state.Wait(true, time.Second*15)

			t.Log(idx, err)
		}()
	}

	<-time.After(time.Second * 5)
	t.Log("5s after, setting flag true")

	state.SetFlag(true)
	t.Log("flag set")
}
