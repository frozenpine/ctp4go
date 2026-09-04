package state_test

import (
	"log/slog"
	"testing"
	"time"

	"github.com/frozenpine/ctp4go/state"
)

type flag uint8

func (v *flag) Migrate(other flag) error {
	*v = other
	return nil
}

func TestState(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug - 2)

	sBool := state.NewBaseFlag[bool, struct{}]("bool")
	sFlag := state.NewBaseFlag[flag, struct{}]("flag")

	for idx := range 10 {
		go func() {
			err := sBool.Wait(true, time.Second*15)

			t.Log(idx, err)
		}()

		go func() {
			err := sFlag.Wait(1, time.Second*15)

			t.Log(idx, err)
		}()
	}

	<-time.After(time.Second * 5)
	t.Log("5s after, setting flag value")

	sBool.SetFlag(true)
	sFlag.SetFlag(1)
	t.Log("flag set")
}
