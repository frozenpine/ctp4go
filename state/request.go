package state

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/frozenpine/ctp4go/thost"
)

var (
	ErrMethodNotFound     = errors.New("method not found")
	ErrMethodArgMissmatch = errors.New("method arg mis-match")
)

type Request interface {
	Type() string
	WithPayload(any) Request

	Do(int64) error
}

type request[API any, D thost.ThostData, PTR DataPtr[D]] struct {
	payload PTR
	api     API
	handler func(API, PTR, int) int
}

func (r *request[API, D, PTR]) Type() string { return r.payload.Type() }

func (r *request[API, D, PTR]) WithPayload(v any) Request {
	if p, ok := v.(PTR); ok {
		return &request[API, D, PTR]{
			payload: p,
			api:     r.api,
			handler: r.handler,
		}
	}

	return nil
}

func (r *request[API, D, DATA]) Do(seq int64) error {
	return thost.Rtn{
		Code: r.handler(r.api, r.payload, int(seq)),
	}.Error()
}

type RequestCache struct {
	lock     sync.RWMutex
	seq      int64
	inflight map[int64]struct{}
}

func (c *RequestCache) DoRequest(r Request) error {
	c.lock.RLock()
	if len(c.inflight) > 0 {
		defer c.lock.RUnlock()

		return errors.New("inflight request execeeded")
	}
	c.lock.RUnlock()

	c.lock.Lock()
	defer c.lock.Unlock()

	seq := c.seq + 1
	c.inflight[seq] = struct{}{}

	return r.Do(seq)
}

func (c *RequestCache) Complete(seq int64) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.inflight, seq)
}

func (c *RequestCache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.seq = 0
	c.inflight = make(map[int64]struct{})
}

type RequestFactory[API any] struct {
	RequestCache

	Api       API
	reqMapper map[reflect.Type]reflect.Method
	preBuilds map[string]Request
}

func NewRequestFactory[API any](api API) *RequestFactory[API] {
	factory := RequestFactory[API]{
		RequestCache: RequestCache{
			inflight: map[int64]struct{}{},
		},
		Api:       api,
		reqMapper: make(map[reflect.Type]reflect.Method),
		preBuilds: make(map[string]Request),
	}

	apiType := reflect.TypeFor[API]()

	for fn := range apiType.Methods() {
		if !strings.HasPrefix(fn.Name, "Req") {
			continue
		}

		dataType := fn.Type.In(0)

		factory.reqMapper[dataType] = fn
	}

	return &factory
}

func (fac *RequestFactory[API]) GetPrebuild(name string) Request {
	fac.lock.RLock()
	defer fac.lock.RUnlock()

	return fac.preBuilds[name]
}

func MakeRequest[
	API any, D thost.ThostData, PTR DataPtr[D],
](factory *RequestFactory[API], v PTR) (Request, error) {
	dataType := reflect.TypeFor[PTR]()

	factory.lock.RLock()
	if req := factory.GetPrebuild(v.Type()); req != nil {
		req = req.WithPayload(v)

		if req != nil {
			defer factory.lock.RUnlock()
			return req, nil
		}
	}

	fn, ok := factory.reqMapper[dataType]
	if !ok {
		return nil, fmt.Errorf(
			"%w: no request method for %+v",
			ErrMethodNotFound, dataType,
		)
	}

	factory.lock.RUnlock()

	req := request[API, D, PTR]{api: factory.Api, payload: v}

	reflect.ValueOf(&req.handler).Elem().Set(reflect.MakeFunc(
		reflect.TypeOf(req.handler),
		func(args []reflect.Value) (results []reflect.Value) {
			fnImpl := args[0].Method(fn.Index)
			return fnImpl.Call(args[1:])
		},
	))

	factory.lock.Lock()
	defer factory.lock.Unlock()

	factory.preBuilds[v.Type()] = &req

	return &req, nil
}
