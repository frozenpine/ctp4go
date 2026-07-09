package state

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"maps"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/frozenpine/ctp4go/thost"
)

const (
	DEFAULT_INFLIGHT_REQUESTS = 1
	DEFAULT_RETRY_SUSPEND     = time.Second
)

var (
	ErrMethodNotFound     = errors.New("method not found")
	ErrMethodArgMissmatch = errors.New("method arg mis-match")

	ErrInflightNotFound = errors.New("inflight not found")
)

type Request interface {
	Type() string
	Executor() string
	WithPayload(any) Request

	Execute(int) error
}

func GetRequestData[
	API any, D thost.ThostData, PTR DataPtr[D],
](r Request) (PTR, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}

	v, ok := r.(*request[API, D, PTR])
	if !ok {
		return nil, errors.New("unsupported request")
	}

	return v.payload, nil
}

type request[API any, D thost.ThostData, PTR DataPtr[D]] struct {
	fnName  string
	payload PTR
	api     API
	handler func(API, PTR, int) int
}

func (r *request[API, D, PTR]) Type() string { return r.payload.Type() }

func (r *request[API, D, PTR]) Executor() string { return r.fnName }

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

func (r *request[API, D, DATA]) Execute(seq int) error {
	slog.Info(
		"start to execute request",
		slog.Int("seq", seq),
		slog.String("executor", r.fnName),
		slog.Any("payload", r.payload),
	)

	return thost.Rtn{
		Code: r.handler(r.api, r.payload, seq),
	}.Error()
}

type reqWait struct {
	ctx    context.Context
	cancel context.CancelFunc
	err    error
}

func (w *reqWait) Wait(timeout time.Duration) error {
	waitCtx := context.Background()

	if timeout > 0 {
		var waitCancel context.CancelFunc

		waitCtx, waitCancel = context.WithCancel(w.ctx)
		defer waitCancel()
	}

	select {
	case <-w.ctx.Done():
		return w.err
	case <-waitCtx.Done():
		return waitCtx.Err()
	}
}

type RequestCache struct {
	lock sync.RWMutex

	reqRootCtx context.Context
	inflights  map[int]*reqWait
	cache      []Request
}

type reqCfg struct {
	inflight int
	attempts int
	maxFail  int
	suspend  time.Duration
}

type reqOpt func(*reqCfg) error

type ReqOptions []reqOpt

func (c *RequestCache) DoRequestAndWait(
	r Request, options ...reqOpt,
) (*reqWait, error) {
	cfg := reqCfg{
		inflight: DEFAULT_INFLIGHT_REQUESTS,
		attempts: 1,
		suspend:  DEFAULT_RETRY_SUSPEND,
	}

	for _, opt := range options {
		if opt == nil {
			continue
		}

		if err := opt(&cfg); err != nil {
			return nil, err
		}
	}

	c.lock.RLock()
	if len(c.inflights) >= cfg.inflight {
		defer c.lock.RUnlock()

		return nil, errors.New("inflight request execeeded")
	}
	c.lock.RUnlock()

	c.lock.Lock()
	defer c.lock.Unlock()

	var (
		failCount int
		err       error
	)

	for range cfg.attempts {
		if failCount > cfg.maxFail {
			break
		}

		c.cache = append(c.cache, r)
		seq := len(c.cache)

		if err = r.Execute(seq); err == nil {
			var wait reqWait
			context.WithCancelCause(c.reqRootCtx)
			wait.ctx, wait.cancel = context.WithCancel(c.reqRootCtx)
			c.inflights[seq] = &wait
			return &wait, nil
		}

		slog.Error(
			"request execute failed",
			slog.Any("error", err),
			slog.Int("seq", seq),
			slog.String("executor", r.Executor()),
		)
		failCount++

		<-time.After(cfg.suspend)
	}

	return nil, err
}

func (c *RequestCache) DoRequest(r Request, options ...reqOpt) error {
	_, err := c.DoRequestAndWait(r, options...)
	return err
}

func (c *RequestCache) Wait(seq int, timeout time.Duration) error {
	if wait := c.getInflight(seq); wait != nil {
		return wait.Wait(timeout)
	}

	return fmt.Errorf("%w: request[%d] completed", ErrInflightNotFound, seq)
}

func (c *RequestCache) getInflight(seq int) *reqWait {
	c.lock.RLock()
	defer c.lock.RUnlock()

	return c.inflights[seq]
}

func (c *RequestCache) Complete(seq int, err error) {
	if wait := c.getInflight(seq); wait != nil {
		wait.err = err
		wait.cancel()
	} else {
		return
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.inflights, seq)
}

func (c *RequestCache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.inflights = make(map[int]*reqWait)
	c.cache = c.cache[:0]
}

type RequestFactory[API any] struct {
	RequestCache

	Api       API
	reqMapper map[reflect.Type]reflect.Method
	preBuilds map[string]Request
}

func NewRequestFactory[API any](ctx context.Context, api API) *RequestFactory[API] {
	factory := RequestFactory[API]{
		RequestCache: RequestCache{
			inflights: map[int]*reqWait{},
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

func (fac *RequestFactory[API]) FilterReqFn(
	filter func(reflect.Method) bool,
) []reflect.Method {
	if filter == nil {
		return nil
	}

	results := make([]reflect.Method, 0, len(fac.reqMapper))

	for fn := range maps.Values(fac.reqMapper) {
		if filter(fn) {
			results = append(results, fn)
		}
	}

	return results
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

	req := request[API, D, PTR]{api: factory.Api, payload: v, fnName: fn.Name}

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
