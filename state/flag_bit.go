package state

import (
	"context"
	"log/slog"
	"reflect"
	"weak"
)

type BitValue interface {
	~int | ~uint | ~int8 | ~uint8 |
		~int16 | ~uint16 | ~int32 | ~uint32 |
		~int64 | ~uint64
}

type BitFlag[T BitValue, V any] struct {
	BaseFlag[T, V]
}

func NewBitFlag[T BitValue, V any](
	name string, options ...flagInitOpt[T, V],
) (flag *BitFlag[T, V]) {
	flagType := reflect.TypeFor[T]()

	if flagType.Kind() != reflect.Pointer {
		flagType = reflect.PointerTo(flagType)
	}

	flag = &BitFlag[T, V]{
		BaseFlag: BaseFlag[T, V]{
			name:     name,
			notifies: map[weak.Pointer[notify]]struct{}{},
		},
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		opt(&flag.BaseFlag)
	}

	if fn, ok := flagType.MethodByName("Migrate"); ok {
		flag.migrator = fn.Func.Interface().(func(*T, T) error)
	}

	return
}

func (b *BitFlag[T, V]) SetFlag(v T, payload ...V) error {
	b.Lock()
	defer b.Unlock()

	if b.migrator != nil {
		err := b.migrator(&b.flag, v)
		if err != nil {
			return err
		}
	} else {
		b.flag |= v
	}

	if len(payload) > 0 {
		b.payload = payload[0]
	}

	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"state flag setted, find waiting",
		slog.String("name", b.name),
		slog.Any("flag", v),
	)

	b.notifyAll()
	return nil
}

func (b *BitFlag[T, V]) CheckFlag(v T) bool {
	b.RLock()
	defer b.RUnlock()

	return b.flag&v == v
}
