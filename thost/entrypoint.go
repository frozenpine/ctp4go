package thost

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/frozenpine/ctp4go/thost/types"
)

const (
	DEFAULT_FLOW_PATH = "./flow"
	DEFAULT_FLOW_MODE = types.THOST_TERT_QUICK
)

type paramKey string

var (
	// ParamRunMode true for Product, false for Test
	ParamRunMode paramKey = "runMode"

	ErrInvalidCreator = errors.New("invalid creator")
)

type Param struct {
	Key   paramKey
	Value any
}

type TraderMaker func(
	libPath, flowPath string, params ...Param,
) func() (TraderApi, error)

type traderMaker struct {
	TraderMaker
	version string
}

func (u *traderMaker) GetVersionTag() string {
	return u.version
}

var (
	tdMaker atomic.Pointer[traderMaker]

	ErrCreatorMissing  = errors.New("thost api creator missing")
	ErrCreatorConflict = errors.New(
		"multiple version's thost api creator conflict",
	)
)

func SetTraderMaker(version string, fn TraderMaker) error {
	if tdMaker.CompareAndSwap(nil, &traderMaker{
		TraderMaker: fn,
		version:     version,
	}) {
		return nil
	}

	exist := tdMaker.Load()

	return fmt.Errorf(
		"%w: trader api[%s] conflicted with [%s]",
		ErrInvalidCreator, version, exist.version,
	)
}

func GetTraderMaker() *traderMaker {
	return tdMaker.Load()
}
