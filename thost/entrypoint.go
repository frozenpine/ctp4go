package thost

import (
	"fmt"
	"sync"
)

type paramKey string

var (
	// ParamFlowPath thost flow dir
	ParamFlowPath paramKey = "flowPath"

	// ParamIsProductionMode true for Product, false for Test
	ParamIsProductionMode paramKey = "runMode"

	// ParamIsUsingUdp true for transport in UDP
	ParamIsUsingUdp paramKey = "useUDP"

	// ParamIsMulticast true for transport by Multicast
	ParamIsMulticast paramKey = "useMulticast"
)

type Param struct {
	Key   paramKey
	Value any
}

type SdkMaker[T any] func(
	libPath string, params ...Param,
) func() (T, error)

type Sdk[T any] struct {
	SdkMaker[T]
	version string
}

func (s *Sdk[T]) GetVersionTag() string { return s.version }

var (
	sdkCache sync.Map
)

func SetSdkMaker[T any](
	plat, sdk, ver string, fn SdkMaker[T],
) error {
	sdkKey := fmt.Sprintf("%s.%s", plat, sdk)

	cache, exist := sdkCache.Load(sdkKey)

	if exist {
		if old, ok := cache.(*Sdk[T]); ok {
			return fmt.Errorf(
				"%w: sdk[%s] conflicted with [%s]",
				ErrInvalidCreator, sdkKey, old.version,
			)
		}
	}

	sdkCache.Store(sdkKey, &Sdk[T]{
		SdkMaker: fn,
		version:  ver,
	})

	return nil
}

func GetSdkMaker[T any](plat, sdk string) (*Sdk[T], error) {
	sdkKey := fmt.Sprintf("%s.%s", plat, sdk)

	cache, exist := sdkCache.Load(sdkKey)
	if !exist {
		return nil, fmt.Errorf(
			"%w: sdk[%s] creator missing",
			ErrCreatorMissing, sdkKey,
		)
	}

	if maker, ok := cache.(*Sdk[T]); ok {
		return maker, nil
	}

	sdkCache.Delete(sdkKey)

	return nil, fmt.Errorf(
		"%w: sdk[%s] creator invalid",
		ErrInvalidCreator, sdkKey,
	)
}
