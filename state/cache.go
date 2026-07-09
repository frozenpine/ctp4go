package state

import (
	"errors"
	"fmt"
	"iter"
	"math"
	"reflect"
	"slices"
	"sync"
	"unsafe"

	"github.com/frozenpine/ctp4go"
	"github.com/frozenpine/ctp4go/thost"
)

var (
	ErrNoFieldName      = errors.New("no field name")
	ErrInvalidFieldType = errors.New("invalid field type")
	ErrInvalidOffset    = errors.New("invalid offset")

	ErrCacheEmpty       = errors.New("cache is nil")
	ErrCacheMismatch    = errors.New("cache miss match")
	ErrCacheDataMissing = errors.New("data not found")
)

type Data interface {
	thost.ThostData

	RawPtr() any

	GetFieldString(string) (string, error)
	GetFieldInt(string) (int64, error)
	GetFieldUInt(string) (uint64, error)
	GetFieldFloat(string) (float64, error)
	GetFieldBool(string) (bool, error)
	GetFieldByte(string, ...int) (byte, error)
}

type Cache interface {
	Size() int
	GetByKey(string) (Data, error)
	GetByIdx(int) (Data, error)
	Iter(...func(Data) bool) iter.Seq2[int, Data]
}

type DataPtr[T thost.ThostData] interface {
	ctp4go.PtrConstraint[T]

	thost.ThostData
}

type dataCfg[T thost.ThostData, Ptr DataPtr[T]] struct {
	dataMerger func(dst Ptr, src Ptr) error

	idtKeys map[string]func(Ptr) string
	fields  map[string]reflect.StructField
}

type DataContainer[T thost.ThostData, Ptr DataPtr[T]] struct {
	dataCfg[T, Ptr]

	lock    sync.RWMutex
	data    Ptr
	basePtr uintptr
}

type dataOpt[T thost.ThostData, Ptr DataPtr[T]] func(*dataCfg[T, Ptr]) error

func WithIdentifier[T thost.ThostData, Ptr DataPtr[T]](
	name string, fn func(Ptr) string,
) dataOpt[T, Ptr] {
	return func(wc *dataCfg[T, Ptr]) error {
		if name == "" {
			return errors.New("invalid identifier name")
		}

		if fn == nil {
			return errors.New("invalid identifier func")
		}

		if _, exist := wc.idtKeys[name]; exist {
			return errors.New("identifier duplicated")
		}

		wc.idtKeys[name] = fn

		return nil
	}
}

func WithMerger[T thost.ThostData, Ptr DataPtr[T]](
	fn func(Ptr, Ptr) error,
) dataOpt[T, Ptr] {
	return func(wc *dataCfg[T, Ptr]) error {
		if fn == nil {
			return errors.New("invalid merger func")
		}

		wc.dataMerger = fn
		return nil
	}
}

func ContainerMaker[T thost.ThostData, Ptr DataPtr[T]](
	options ...dataOpt[T, Ptr],
) (*dataCfg[T, Ptr], func(Ptr) *DataContainer[T, Ptr], error) {
	ptrType := reflect.TypeFor[Ptr]()
	dType := ptrType.Elem()

	if dType.Kind() != reflect.Struct {
		return nil, nil, errors.New("generic type must be struct")
	}

	cfg := dataCfg[T, Ptr]{
		dataMerger: func(dst, src Ptr) error {
			*dst = *src
			return nil
		},
		idtKeys: make(map[string]func(Ptr) string),
		fields:  make(map[string]reflect.StructField),
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return nil, nil, err
		}
	}

	for f := range dType.Fields() {
		cfg.fields[f.Name] = f
	}

	return &cfg, func(t Ptr) *DataContainer[T, Ptr] {
		// 封装时copy数据，避免传入指针为栈指针
		var copy T = *t
		return &DataContainer[T, Ptr]{
			dataCfg: cfg,
			data:    &copy,
			basePtr: uintptr(unsafe.Pointer(t)),
		}
	}, nil
}

func (w *DataContainer[T, Ptr]) String() string { return w.data.String() }

func (w *DataContainer[T, Ptr]) Type() string { return w.data.Type() }

func (w *DataContainer[T, Ptr]) Merge(v Ptr) {
	w.lock.Lock()
	defer w.lock.Unlock()

	w.dataMerger(w.data, v)
}

func (w *DataContainer[T, Ptr]) RawPtr() any {
	return w.data
}

func (w *DataContainer[T, Ptr]) Data() Ptr {
	return w.data
}

func (w *DataContainer[T, Ptr]) GetFieldString(name string) (string, error) {
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

func (w *DataContainer[T, Ptr]) GetFieldInt(name string) (int64, error) {
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

func (w *DataContainer[T, Ptr]) GetFieldUInt(name string) (uint64, error) {
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

func (w *DataContainer[T, Ptr]) GetFieldBool(name string) (bool, error) {
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

func (w *DataContainer[T, Ptr]) GetFieldFloat(name string) (float64, error) {
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

func (w *DataContainer[T, Ptr]) GetFieldByte(
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

type DataCache[T thost.ThostData, Ptr DataPtr[T]] struct {
	lock sync.RWMutex

	cfg       *dataCfg[T, Ptr]
	idtCache  map[string]int
	cache     []*DataContainer[T, Ptr]
	dataMaker func(Ptr) *DataContainer[T, Ptr]
}

func NewDataCache[T thost.ThostData, Ptr DataPtr[T]](
	options ...dataOpt[T, Ptr],
) (*DataCache[T, Ptr], error) {
	cfg, maker, err := ContainerMaker(options...)

	if err != nil {
		return nil, err
	}

	return &DataCache[T, Ptr]{
		cfg:       cfg,
		idtCache:  make(map[string]int),
		dataMaker: maker,
	}, nil
}

func (c *DataCache[T, Ptr]) AddOrUpdate(v Ptr) int {
	if v == nil {
		return -1
	}

	data := c.dataMaker(v)

	c.lock.Lock()
	defer c.lock.Unlock()

	var (
		idx    int = -1
		rtnIdx     = len(c.cache)
		merged bool
	)

	for _, identifier := range c.cfg.idtKeys {
		idt := identifier(v)

		var (
			exist  bool
			preIdx = idx
		)

		idx, exist = c.idtCache[idt]

		if exist {
			rtnIdx = idx
			if !merged {
				c.cache[idx].Merge(v)
				merged = true
			}

			if preIdx >= 0 && idx != preIdx {
				panic("exists cache index missmatch")
			}
		} else {
			c.idtCache[idt] = rtnIdx
		}
	}

	if !merged {
		c.cache = append(c.cache, data)
	}

	return rtnIdx
}

func (c *DataCache[T, Ptr]) GetByKey(k string) (*DataContainer[T, Ptr], error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	idx, exist := c.idtCache[k]
	if !exist {
		return nil, fmt.Errorf(
			"%w: invalid key %s", ErrCacheDataMissing, k,
		)
	}

	return c.cache[idx], nil
}

func (c *DataCache[T, Ptr]) GetByIdx(idx int) (*DataContainer[T, Ptr], error) {
	if idx < 0 {
		return nil, fmt.Errorf(
			"%w: index[%d] out of range", ErrCacheDataMissing, idx,
		)
	}

	c.lock.RLock()
	defer c.lock.RUnlock()

	if idx >= len(c.cache) {
		return nil, fmt.Errorf(
			"%w: index[%d] out of range", ErrCacheDataMissing, idx,
		)
	}

	return c.cache[idx], nil
}

func (c *DataCache[T, Ptr]) Size() int {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return len(c.cache)
}

func (c *DataCache[T, Ptr]) Iter(
	filters ...func(*DataContainer[T, Ptr]) bool,
) iter.Seq2[int, *DataContainer[T, Ptr]] {
	shortCircuit := func(v *DataContainer[T, Ptr]) bool {
		for _, f := range filters {
			if !f(v) {
				return false
			}
		}

		return true
	}

	c.lock.RLock()
	defer c.lock.RUnlock()

	snap := slices.Clone(c.cache)

	return func(yield func(int, *DataContainer[T, Ptr]) bool) {
		for idx, v := range snap {
			if !shortCircuit(v) {
				continue
			}

			if !yield(idx, v) {
				return
			}
		}
	}
}

type ReadOnlyCache[T thost.ThostData, Ptr DataPtr[T]] struct {
	cache *DataCache[T, Ptr]
}

func NewReadOnlyCache[T thost.ThostData, Ptr DataPtr[T]](
	c *DataCache[T, Ptr],
) (Cache, error) {
	if c == nil {
		return nil, fmt.Errorf("%w: data cache empty", ErrCacheEmpty)
	}
	return ReadOnlyCache[T, Ptr]{c}, nil
}

func CastDataCache[T thost.ThostData, Ptr DataPtr[T]](
	c Cache,
) (*DataCache[T, Ptr], error) {
	if c == nil {
		return nil, fmt.Errorf("%w: cache interface empty", ErrCacheEmpty)
	}

	if v, ok := c.(*ReadOnlyCache[T, Ptr]); ok {
		return v.cache, nil
	}

	return nil, fmt.Errorf(
		"%w: cache is not a readonly cache", ErrCacheMismatch,
	)
}

func (r ReadOnlyCache[T, Ptr]) Size() int { return r.Size() }

func (r ReadOnlyCache[T, Ptr]) GetByKey(k string) (Data, error) {
	return r.cache.GetByKey(k)
}

func (r ReadOnlyCache[T, Ptr]) GetByIdx(idx int) (Data, error) {
	return r.cache.GetByIdx(idx)
}

func (r ReadOnlyCache[T, Ptr]) Iter(
	filters ...func(Data) bool,
) iter.Seq2[int, Data] {
	bridgeFilters := make([]func(*DataContainer[T, Ptr]) bool, len(filters))
	for idx, fn := range filters {
		bridgeFilters[idx] = func(dc *DataContainer[T, Ptr]) bool {
			return fn(dc)
		}
	}

	return func(yield func(int, Data) bool) {
		for idx, v := range r.cache.Iter(bridgeFilters...) {
			if !yield(idx, v) {
				return
			}
		}
	}
}
