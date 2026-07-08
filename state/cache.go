package state

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"sync"
	"unsafe"

	"github.com/frozenpine/ctp4go"
)

var (
	ErrNoFieldName      = errors.New("no field name")
	ErrInvalidFieldType = errors.New("invalid field type")
	ErrInvalidOffset    = errors.New("invalid offset")
)

type Data interface {
	fmt.Stringer

	GetIdentity() string

	GetFieldString(string) (string, error)
	GetFieldInt(string) (int64, error)
	GetFieldUInt(string) (uint64, error)
	GetFieldFloat(string) (float64, error)
	GetFieldBool(string) (bool, error)
	GetFieldByte(string, ...int) (byte, error)
}

type dataPtr[T fmt.Stringer] interface {
	ctp4go.PtrConstraint[T]

	fmt.Stringer
}

type dataCfg[T fmt.Stringer, Ptr dataPtr[T]] struct {
	identifier func(Ptr) string
	dataMerger func(dst Ptr, src Ptr) error

	fields map[string]reflect.StructField
}

type dataContainer[T fmt.Stringer, Ptr dataPtr[T]] struct {
	dataCfg[T, Ptr]

	lock    sync.RWMutex
	data    Ptr
	basePtr uintptr
}

type wrapperOpt[T fmt.Stringer, Ptr dataPtr[T]] func(*dataCfg[T, Ptr]) error

func WithIdentifier[T fmt.Stringer, Ptr dataPtr[T]](
	fn func(Ptr) string,
) wrapperOpt[T, Ptr] {
	return func(wc *dataCfg[T, Ptr]) error {
		if fn == nil {
			return errors.New("invalid identifier func")
		}

		wc.identifier = fn
		return nil
	}
}

func WithMerger[T fmt.Stringer, Ptr dataPtr[T]](
	fn func(Ptr, Ptr) error,
) wrapperOpt[T, Ptr] {
	return func(wc *dataCfg[T, Ptr]) error {
		if fn == nil {
			return errors.New("invalid merger func")
		}

		wc.dataMerger = fn
		return nil
	}
}

func MakeDataWrapper[T fmt.Stringer, Ptr dataPtr[T]](
	options ...wrapperOpt[T, Ptr],
) (func(Ptr) *dataContainer[T, Ptr], error) {
	ptrType := reflect.TypeFor[Ptr]()
	dType := ptrType.Elem()

	if dType.Kind() != reflect.Struct {
		return nil, errors.New("generic type must be struct")
	}

	cfg := dataCfg[T, Ptr]{
		fields: make(map[string]reflect.StructField),
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	if fn, ok := dType.MethodByName("GetIdentity"); ok {
		reflect.ValueOf(&cfg.identifier).Elem().Set(reflect.MakeFunc(
			reflect.TypeOf(cfg.identifier),
			func(args []reflect.Value) (results []reflect.Value) {
				return fn.Func.Call(args)
			},
		))
	}

	if fn, ok := ptrType.MethodByName("Merge"); ok {
		reflect.ValueOf(&cfg.identifier).Elem().Set(reflect.MakeFunc(
			reflect.TypeOf(cfg.identifier),
			func(args []reflect.Value) (results []reflect.Value) {
				return fn.Func.Call(args)
			},
		))
	}

	if cfg.identifier == nil || cfg.dataMerger == nil {
		return nil, errors.New("no identifier & merger specified")
	}

	for f := range dType.Fields() {
		cfg.fields[f.Name] = f
	}

	return func(t Ptr) *dataContainer[T, Ptr] {
		return &dataContainer[T, Ptr]{
			dataCfg: cfg,
			data:    t,
			basePtr: uintptr(unsafe.Pointer(t)),
		}
	}, nil
}

func (w *dataContainer[T, Ptr]) String() string {
	return w.data.String()
}

func (w *dataContainer[T, Ptr]) GetIdentity() string {
	w.lock.RLock()
	defer w.lock.RUnlock()

	return w.identifier(w.data)
}

func (w *dataContainer[T, Ptr]) Merge(v Ptr) {
	w.lock.Lock()
	defer w.lock.Unlock()

	w.dataMerger(w.data, v)
}

func (w *dataContainer[T, Ptr]) GetFieldString(name string) (string, error) {
	f, exist := w.fields[name]
	if !exist {
		return "", fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.String, reflect.Slice:
		return unsafe.String(
			(*byte)(unsafe.Pointer(*(*uintptr)(unsafe.Pointer(ptr)))),
			*(*int)(unsafe.Pointer(ptr + 8)),
		), nil
	case reflect.Array:
		return ctp4go.DecodeGBK(unsafe.Slice(
			(*byte)(unsafe.Pointer(ptr)), f.Type.Size(),
		)), nil
	default:
		return "", fmt.Errorf(
			"%w: cannot convert %s to string",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

func (w *dataContainer[T, Ptr]) GetFieldInt(name string) (int64, error) {
	f, exist := w.fields[name]
	if !exist {
		return 0, fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.Int32:
		return int64(*(*int32)(unsafe.Pointer(ptr))), nil
	case reflect.Uint32:
		return int64(*(*uint32)(unsafe.Pointer(ptr))), nil
	case reflect.Int16:
		return int64(*(*int16)(unsafe.Pointer(ptr))), nil
	case reflect.Uint16:
		return int64(*(*uint16)(unsafe.Pointer(ptr))), nil
	case reflect.Int8:
		return int64(*(*int8)(unsafe.Pointer(ptr))), nil
	case reflect.Uint8:
		return int64(*(*uint8)(unsafe.Pointer(ptr))), nil
	case reflect.Int64:
		return *(*int64)(unsafe.Pointer(ptr)), nil
	case reflect.Int:
		return int64(*(*int)(unsafe.Pointer(ptr))), nil
	default:
		return 0, fmt.Errorf(
			"%w: cannot convert %s to int64",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

func (w *dataContainer[T, Ptr]) GetFieldUInt(name string) (uint64, error) {
	f, exist := w.fields[name]
	if !exist {
		return 0, fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.Uint32:
		return uint64(*(*uint32)(unsafe.Pointer(ptr))), nil
	case reflect.Uint16:
		return uint64(*(*uint16)(unsafe.Pointer(ptr))), nil
	case reflect.Uint8:
		return uint64(*(*uint8)(unsafe.Pointer(ptr))), nil
	case reflect.Uint64:
		return *(*uint64)(unsafe.Pointer(ptr)), nil
	case reflect.Uint:
		return uint64(*(*uint)(unsafe.Pointer(ptr))), nil
	default:
		return 0, fmt.Errorf(
			"%w: cannot convert %s to uint64",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

func (w *dataContainer[T, Ptr]) GetFieldBool(name string) (bool, error) {
	f, exist := w.fields[name]
	if !exist {
		return false, fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.Int32:
		return *(*int32)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Uint32:
		return *(*uint32)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Int16:
		return *(*int16)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Uint16:
		return *(*uint16)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Int8:
		return *(*int8)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Uint8:
		return *(*uint8)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Int64:
		return *(*int64)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Uint64:
		return *(*uint64)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Int:
		return *(*int)(unsafe.Pointer(ptr)) != 0, nil
	case reflect.Uint:
		return *(*uint)(unsafe.Pointer(ptr)) != 0, nil
	default:
		return false, fmt.Errorf(
			"%w: cannot convert %s to int64",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

func (w *dataContainer[T, Ptr]) GetFieldFloat(name string) (float64, error) {
	f, exist := w.fields[name]
	if !exist {
		return math.NaN(), fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.Int32:
		return float64(*(*int32)(unsafe.Pointer(ptr))), nil
	case reflect.Uint32:
		return float64(*(*uint32)(unsafe.Pointer(ptr))), nil
	case reflect.Int16:
		return float64(*(*int16)(unsafe.Pointer(ptr))), nil
	case reflect.Uint16:
		return float64(*(*uint16)(unsafe.Pointer(ptr))), nil
	case reflect.Int8:
		return float64(*(*int8)(unsafe.Pointer(ptr))), nil
	case reflect.Uint8:
		return float64(*(*uint8)(unsafe.Pointer(ptr))), nil
	case reflect.Int64:
		return float64(*(*int64)(unsafe.Pointer(ptr))), nil
	case reflect.Int:
		return float64(*(*int)(unsafe.Pointer(ptr))), nil
	case reflect.Float64:
		return *(*float64)(unsafe.Pointer(ptr)), nil
	case reflect.Float32:
		return float64(*(*float32)(unsafe.Pointer(ptr))), nil
	default:
		return math.NaN(), fmt.Errorf(
			"%w: cannot convert %s to float64",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

func (w *dataContainer[T, Ptr]) GetFieldByte(
	name string, offset ...int,
) (byte, error) {
	f, exist := w.fields[name]
	if !exist {
		return 0, fmt.Errorf("%w: %s", ErrNoFieldName, name)
	}

	ptr := w.basePtr + f.Offset
	pos := 0
	if len(offset) > 0 {
		if offset[0] < 0 {
			return 0, fmt.Errorf(
				"%w: offset[%d] out of range", ErrInvalidOffset, offset[0],
			)
		}

		pos = offset[0]
	}
	w.lock.RLock()
	defer w.lock.RUnlock()

	switch f.Type.Kind() {
	case reflect.Uint8, reflect.Int8:
		return *(*byte)(unsafe.Pointer(ptr)), nil
	case reflect.Array:
		if pos >= int(f.Type.Size()) {
			return 0, fmt.Errorf(
				"%w: offset[%d] out of range", ErrInvalidOffset, offset[0],
			)
		}

		return *(*byte)(unsafe.Pointer(ptr + uintptr(pos))), nil
	case reflect.String:
		len := *(*int)(unsafe.Pointer(ptr + 8))

		if pos >= len {
			return 0, fmt.Errorf(
				"%w: offset[%d] out of range", ErrInvalidOffset, offset[0],
			)
		}

		return *(*byte)(unsafe.Pointer(
			*(*uintptr)(unsafe.Pointer(ptr)) + uintptr(pos),
		)), nil
	default:
		return 0, fmt.Errorf(
			"%w: cannot convert %s to byte",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

type TypedCache[T fmt.Stringer, Ptr dataPtr[T]] struct {
	lock sync.RWMutex

	idtCache  map[string]int
	cache     []*dataContainer[T, Ptr]
	dataMaker func(Ptr) *dataContainer[T, Ptr]
}

func (c *TypedCache[T, Ptr]) AddOrUpdate(v Ptr) {
	c.lock.Lock()
	defer c.lock.Unlock()

	data := c.dataMaker(v)
	idt := data.GetIdentity()

	if idx, exist := c.idtCache[idt]; exist {
		c.cache[idx].Merge(v)
	} else {
		idx := len(c.idtCache)
		c.idtCache[idt] = idx
		c.cache = append(c.cache, data)
	}
}

func (c *TypedCache[T, Ptr]) Get(k string) (*dataContainer[T, Ptr], bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	idx, exist := c.idtCache[k]
	if !exist {
		return nil, false
	}

	return c.cache[idx], true
}
