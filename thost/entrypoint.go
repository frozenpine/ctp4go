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
	// ParamFlowPath thost flow dir
	ParamFlowPath paramKey = "flowPath"

	// ParamRunMode true for Product, false for Test
	ParamRunMode paramKey = "runMode"

	// ParamUseUDP true for transport in UDP
	ParamUseUDP paramKey = "useUDP"

	// ParamUseMulticast true for transport by Multicast
	ParamUseMulticast paramKey = "useMulticast"

	ErrInvalidCreator = errors.New("invalid creator")
)

type Param struct {
	Key   paramKey
	Value any
}

type TraderMaker func(
	libPath string, params ...Param,
) func() (TraderApi, error)

type traderMaker struct {
	TraderMaker
	version string
}

func (u *traderMaker) GetVersionTag() string {
	return u.version
}

type MduserMaker func(
	libPath string, params ...Param,
) func() (MdApi, error)

type mduserMaker struct {
	MduserMaker
	version string
}

func (u *mduserMaker) GetVersionTag() string {
	return u.version
}

var (
	trader atomic.Pointer[traderMaker]
	mduser atomic.Pointer[mduserMaker]

	ErrCreatorMissing  = errors.New("thost api creator missing")
	ErrCreatorConflict = errors.New(
		"multiple version's thost api creator conflict",
	)
)

func SetTraderMaker(version string, fn TraderMaker) error {
	if trader.CompareAndSwap(nil, &traderMaker{
		TraderMaker: fn,
		version:     version,
	}) {
		return nil
	}

	exist := trader.Load()

	return fmt.Errorf(
		"%w: trader api[%s] conflicted with [%s]",
		ErrInvalidCreator, version, exist.version,
	)
}

func GetTraderMaker() *traderMaker {
	return trader.Load()
}

func SetMduserMaker(version string, fn MduserMaker) error {
	if mduser.CompareAndSwap(nil, &mduserMaker{
		MduserMaker: fn,
		version:     version,
	}) {
		return nil
	}

	exist := trader.Load()

	return fmt.Errorf(
		"%w mduser api[%s] conflicted with [%s]",
		ErrInvalidCreator, version, exist.version,
	)
}

func GetMduserMaker() *mduserMaker {
	return mduser.Load()
}
