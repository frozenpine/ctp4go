package trader

import (
	"testing"

	"github.com/frozenpine/ctp4go/state"
)

func TestStateReflect(t *testing.T) {
	flag := state.NewFlag[traderState]("test")

	t.Log(flag.SetFlag(Connected))

	t.Log(flag.SetFlag(Initialized))
	t.Log(flag.SetFlag(Connected))

	t.Log(flag.GetFlag())
}
