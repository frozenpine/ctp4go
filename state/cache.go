package state

import (
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"log/slog"
	"math"
	"reflect"
	"slices"
	"sync"
	"unsafe"

	"github.com/frozenpine/ctp4go"
)

var (
	ErrNoFieldName      = errors.New("no field name")
	ErrInvalidFieldType = errors.New("invalid field type")
	ErrInvalidOffset    = errors.New("invalid offset")

	ErrCacheEmpty        = errors.New("cache is nil")
	ErrCacheMismatch     = errors.New("cache miss match")
	ErrCacheDataMissing  = errors.New("data not found")
	ErrCacheDataMismatch = errors.New("cache merge data mismatch")
)

type Data interface {
	ctp4go.DataConstraint

	Identities() []string
	Groups() []string

	GetFieldString(name string) (string, error)
	GetFieldInt(name string) (int64, error)
	GetFieldUInt(name string) (uint64, error)
	GetFieldFloat(name string) (float64, error)
	GetFieldBool(name string) (bool, error)
	// GetFieldByte 获取字段的指定字节, 默认获取0偏移字节
	GetFieldByte(name string, offset ...int) (byte, error)
}

func MustGetString(v Data, name string) string {
	d, err := v.GetFieldString(name)
	if err != nil {
		panic(err)
	}
	return d
}

func MustGetInt(v Data, name string) int64 {
	d, err := v.GetFieldInt(name)
	if err != nil {
		panic(err)
	}
	return d
}

func MustGetUInt(v Data, name string) uint64 {
	d, err := v.GetFieldUInt(name)
	if err != nil {
		panic(err)
	}
	return d
}

func MustGetFloat(v Data, name string) float64 {
	d, err := v.GetFieldFloat(name)
	if err != nil {
		panic(err)
	}
	return d
}

func MustGetBool(v Data, name string) bool {
	d, err := v.GetFieldBool(name)
	if err != nil {
		panic(err)
	}
	return d
}

func MustGetByte(v Data, name string, offset ...int) byte {
	d, err := v.GetFieldByte(name, offset...)
	if err != nil {
		panic(err)
	}
	return d
}

type DataHandler[T ctp4go.DataConstraint, Ptr DataPtr[T]] interface {
	WithData(func(Ptr))
}

type Cache interface {
	Name() string
	Size() int
	GetDataByKey(idt string) (Data, error)
	GetDataByIdx(idx int) (Data, error)
	Iter(filters ...func(Data) bool) iter.Seq2[int, Data]

	AddNotifier(func(int, Data)) error
}

type DataPtr[T ctp4go.DataConstraint] interface {
	ctp4go.PtrConstraint[T]

	ctp4go.DataConstraint
}

type dataCfg[T ctp4go.DataConstraint, Ptr DataPtr[T]] struct {
	dataMerger func(dst Ptr, src Ptr) error

	dataName  string
	idtKeys   map[string]func(Ptr) string
	groupKeys map[string]func(Ptr) string
	fieldList []reflect.StructField
	fields    map[string]*reflect.StructField
	fieldFmt  map[string]func(io.Writer, uintptr)

	notifies []func(int, Data)
}

type DataContainer[T ctp4go.DataConstraint, Ptr DataPtr[T]] struct {
	*dataCfg[T, Ptr]

	lock    sync.RWMutex
	data    Ptr
	basePtr uintptr
}

type DataOpt[T ctp4go.DataConstraint, Ptr DataPtr[T]] func(*dataCfg[T, Ptr]) error

type DataOptions[T ctp4go.DataConstraint, Ptr DataPtr[T]] []DataOpt[T, Ptr]

func WithIdentifier[T ctp4go.DataConstraint, Ptr DataPtr[T]](
	name string, fn func(Ptr) string,
) DataOpt[T, Ptr] {
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

func WithMerger[T ctp4go.DataConstraint, Ptr DataPtr[T]](
	fn func(Ptr, Ptr) error,
) DataOpt[T, Ptr] {
	return func(wc *dataCfg[T, Ptr]) error {
		if fn == nil {
			return errors.New("invalid merger func")
		}

		wc.dataMerger = fn
		return nil
	}
}

func WithNotifier[T ctp4go.DataConstraint, Ptr DataPtr[T]](
	fn func(int, Data),
) DataOpt[T, Ptr] {
	return func(dc *dataCfg[T, Ptr]) error {
		if fn == nil {
			return errors.New("invalid data notifier")
		}

		dc.notifies = append(dc.notifies, fn)
		return nil
	}
}

func (w *DataContainer[T, Ptr]) Type() string { return w.data.Type() }

func (w *DataContainer[T, Ptr]) String() string {
	buff := ctp4go.GetStringBuilder()
	fmt.Fprintf(buff, "%s{", w.dataCfg.dataName)
	for idx, f := range w.dataCfg.fieldList {
		stringer := w.dataCfg.fieldFmt[f.Name]
		if idx > 0 {
			buff.WriteString(", ")
		}
		buff.WriteString(f.Name)
		buff.WriteByte('=')
		stringer(buff, w.basePtr)
	}
	buff.WriteByte('}')
	return buff.String()
}

func (w *DataContainer[T, Ptr]) Merge(v Ptr) {
	w.lock.Lock()
	defer w.lock.Unlock()

	w.dataMerger(w.data, v)
}

func (w *DataContainer[T, Ptr]) WithData(fn func(Ptr)) {
	if fn == nil {
		return
	}

	w.lock.Lock()
	defer w.lock.Unlock()

	fn(w.data)
}

func (w *DataContainer[T, Ptr]) Identities() []string {
	result := make([]string, 0, len(w.idtKeys))

	buff := ctp4go.GetStringBuilder()

	for name, idt := range w.idtKeys {
		buff.Reset()
		fmt.Fprintf(buff, "%s:%s", name, idt(w.data))
		result = append(result, buff.String())
	}
	return result
}

func (w *DataContainer[T, Ptr]) Groups() []string {
	result := make([]string, 0, len(w.groupKeys))
	for name, idt := range w.groupKeys {
		result = append(result, name+":"+idt(w.data))
	}
	return result
}

// GetFieldString 获取字段string值
// 支持字段类型: 字符串, 字节数组，字节切片
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

// GetFieldInt 获取字段int64值
// 支持字段类型：所有符号整形，64位以下无符号整型
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

// GetFieldUInt 获取字段uint64值
// 支持字段类型：所有无符号整形
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

// GetFieldBool 获取字段bool值
// 支持字段类型：bool, 所有整形非0为真
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
	case reflect.Bool:
		return *(*bool)(unsafe.Pointer(ptr)), nil
	default:
		return false, fmt.Errorf(
			"%w: cannot convert %s to int64",
			ErrInvalidFieldType, f.Type.Kind(),
		)
	}
}

// GetFieldFloat 获取字段float64值
// 支持所有浮点型，有符号整型，64位以下无符号整型
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

// GetFieldByte 获取字段字节值，默认获取0偏移位置字节
// 支持字段类型：字节型，字节数组，字符串
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

type DataCache[T ctp4go.DataConstraint, Ptr DataPtr[T]] struct {
	lock sync.RWMutex

	cfg         *dataCfg[T, Ptr]
	name        string
	idtPrefixes []string
	idtCache    map[string]int
	grpCache    map[string][]int
	cache       []*DataContainer[T, Ptr]
	dataMaker   func(Ptr) *DataContainer[T, Ptr]
}

func NewDataCache[T ctp4go.DataConstraint, Ptr DataPtr[T]](
	name string, size uint32, options ...DataOpt[T, Ptr],
) (*DataCache[T, Ptr], error) {
	if name == "" {
		return nil, errors.New("cache name empty")
	}

	ptrType := reflect.TypeFor[Ptr]()
	dType := ptrType.Elem()

	if dType.Kind() != reflect.Struct {
		return nil, errors.New("generic type must be struct")
	}

	cfg := dataCfg[T, Ptr]{
		dataMerger: func(dst, src Ptr) error {
			*dst = *src
			return nil
		},
		dataName: dType.Name(),
		idtKeys:  make(map[string]func(Ptr) string),
		fields:   make(map[string]*reflect.StructField),
		fieldFmt: make(map[string]func(io.Writer, uintptr)),
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	for f := range dType.Fields() {
		cfg.fieldList = append(cfg.fieldList, f)
		cfg.fields[f.Name] = &f
		cfg.fieldFmt[f.Name] = func(buff io.Writer, b uintptr) {
			v := reflect.NewAt(f.Type, unsafe.Pointer(b+f.Offset))
			fmt.Fprintf(buff, "%+v", v.Elem().Interface())
		}
	}

	return &DataCache[T, Ptr]{
		cfg:      &cfg,
		idtCache: make(map[string]int),
		grpCache: make(map[string][]int),
		cache: make(
			[]*DataContainer[T, Ptr], 0, ctp4go.NextPowerOfTwo(size),
		),

		dataMaker: func(t Ptr) *DataContainer[T, Ptr] {
			if t == nil {
				return nil
			}

			// 封装时copy数据，避免传入指针为栈指针
			var copy T = *t
			return &DataContainer[T, Ptr]{
				dataCfg: &cfg,
				data:    &copy,
				basePtr: uintptr(unsafe.Pointer(&copy)),
			}
		},
	}, nil
}

func (c *DataCache[T, Ptr]) Name() string { return c.name }

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

	identities := data.Identities()
	for _, idt := range identities {
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
				data = c.cache[idx]

				slog.Log(
					context.Background(), slog.LevelDebug-1,
					"cache data merged",
					slog.String("idt", idt),
					slog.Int("idx", idx),
					slog.Int("size", len(c.cache)),
				)
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
		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"cache data appended",
			slog.Any("identities", identities),
			slog.Int("idx", rtnIdx),
			slog.Int("size", len(c.cache)),
		)
	}

	for _, n := range c.cfg.notifies {
		n(rtnIdx, data)
	}

	return rtnIdx
}

func (c *DataCache[T, Ptr]) GetDataByKey(
	k string,
) (*DataContainer[T, Ptr], error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	key := ctp4go.GetStringBuilder()

	for prefix := range c.cfg.idtKeys {
		fmt.Fprintf(key, "%s:%s", prefix, k)
		idx, exist := c.idtCache[key.String()]

		if exist {
			return c.cache[idx], nil
		}
		key.Reset()
	}

	return nil, fmt.Errorf(
		"%w: invalid key %s", ErrCacheDataMissing, k,
	)
}

func (c *DataCache[T, Ptr]) GetDataByIdx(
	idx int,
) (*DataContainer[T, Ptr], error) {
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

func (c *DataCache[T, Ptr]) Array() []Ptr {
	c.lock.RLock()
	defer c.lock.RUnlock()

	rtn := make([]Ptr, len(c.cache))
	for idx, d := range c.cache {
		rtn[idx] = d.data
	}

	return rtn
}

func (c *DataCache[T, Ptr]) Empty() bool {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return len(c.cache) < 1
}

func (c *DataCache[T, Ptr]) Overflow() bool { return false }

func (c *DataCache[T, Ptr]) First() Ptr {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(c.cache) > 0 {
		return c.cache[0].data
	}

	return nil
}

func (c *DataCache[T, Ptr]) Last() Ptr {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(c.cache) > 0 {
		return c.cache[len(c.cache)-1].data
	}

	return nil
}

func (c *DataCache[T, Ptr]) PushValue(v Ptr) {
	c.AddOrUpdate(v)
}

func (c *DataCache[T, Ptr]) AddNotifier(fn func(int, Data)) error {
	if fn == nil {
		return errors.New("invalid notifier")
	}
	c.lock.Lock()
	defer c.lock.Unlock()

	c.cfg.notifies = append(c.cfg.notifies, fn)
	return nil
}

// ReadOnlyCache 封装DataCache，并抹除数据相关接口的泛型化类型
type ReadOnlyCache[T ctp4go.DataConstraint, Ptr DataPtr[T]] struct {
	*DataCache[T, Ptr]
}

func NewReadOnlyCache[
	T ctp4go.DataConstraint, Ptr DataPtr[T],
](c *DataCache[T, Ptr]) (Cache, error) {
	if c == nil {
		return nil, fmt.Errorf("%w: data cache empty", ErrCacheEmpty)
	}
	return &ReadOnlyCache[T, Ptr]{c}, nil
}

func (r ReadOnlyCache[T, Ptr]) GetDataByKey(k string) (Data, error) {
	return r.DataCache.GetDataByKey(k)
}

func (r ReadOnlyCache[T, Ptr]) GetDataByIdx(idx int) (Data, error) {
	return r.DataCache.GetDataByIdx(idx)
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
		for idx, v := range r.DataCache.Iter(bridgeFilters...) {
			if !yield(idx, v) {
				return
			}
		}
	}
}

// ConvertableRdCache 功能同ReadonlyCache
// 支持Rsp和Rtn同内存布局但不同数据类型签名的指针强制转换
type ConvertableRdCache[
	T ctp4go.DataConstraint, Ptr DataPtr[T],
	CVT ctp4go.DataConstraint, SRC DataPtr[CVT],
] struct{ ReadOnlyCache[T, Ptr] }

func NewConvertableRdCache[
	CVT ctp4go.DataConstraint, SRC DataPtr[CVT],
	T ctp4go.DataConstraint, Ptr DataPtr[T],
](c *DataCache[T, Ptr]) (Cache, error) {
	if c == nil {
		return nil, fmt.Errorf("%w: data cache empty", ErrCacheEmpty)
	}
	return &ConvertableRdCache[T, Ptr, CVT, SRC]{
		ReadOnlyCache[T, Ptr]{c},
	}, nil
}

func CastReadonlyCache[T ctp4go.DataConstraint, Ptr DataPtr[T]](
	c Cache,
) (*DataCache[T, Ptr], error) {
	if c == nil {
		return nil, fmt.Errorf("%w: cache interface empty", ErrCacheEmpty)
	}

	if v, ok := c.(*ReadOnlyCache[T, Ptr]); ok {
		return v.DataCache, nil
	}

	return nil, fmt.Errorf(
		"%w: cache is not a readonly cache", ErrCacheMismatch,
	)
}

func CastConvertableRdCache[
	T ctp4go.DataConstraint, Ptr DataPtr[T],
	CVT ctp4go.DataConstraint, SRC DataPtr[CVT],
](c Cache) (*DataCache[T, Ptr], error) {
	if c == nil {
		return nil, fmt.Errorf("%w: cache interface empty", ErrCacheEmpty)
	}

	if v, ok := c.(*ConvertableRdCache[T, Ptr, CVT, SRC]); ok {
		return v.DataCache, nil
	}

	return nil, fmt.Errorf(
		"%w: cache is not a convertable readonly cache", ErrCacheMismatch,
	)
}
