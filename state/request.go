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
	DEFAULT_QRY_INFLIGHT_REQUESTS = 1

	DEFAULT_FAIL_ATTEMPTS = 1
	DEFAULT_RETRY_SUSPEND = time.Second

	DEFAULT_QRY_REQ_PREFIX = "ReqQry"
)

var (
	ErrMethodNotFound     = errors.New("method not found")
	ErrMethodArgMissmatch = errors.New("method arg mis-match")

	ErrInflightNotFound = errors.New("inflight not found")
)

type Request interface {
	fmt.Stringer

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
	api     API
	handler func(API, PTR, int) int
	payload PTR
}

func (r *request[API, D, PTR]) Type() string { return r.payload.Type() }

func (r *request[API, D, PTR]) Executor() string { return r.fnName }

func (r *request[API, D, PTR]) WithPayload(v any) Request {
	if p, ok := v.(PTR); ok {
		newReq := *r
		newReq.payload = p
		return &newReq
	}

	slog.Error(
		"payload mismatch with request",
		slog.Any("req_type", reflect.TypeFor[PTR]()),
		slog.Any("payload_type", reflect.TypeOf(v)),
	)

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

func (r *request[API, D, PTR]) String() string {
	builder := strings.Builder{}

	fmt.Fprintf(
		&builder, "Request{Executor=%s, Payload=%+v}",
		r.fnName, r.payload,
	)

	return builder.String()
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
	lock       sync.RWMutex
	reqRootCtx context.Context

	cache     []Request
	inflights map[int]*reqWait

	lastQry          time.Time
	qryInflightCount int
	qryLimit         int
	qryLimitCount    int
	qryInflightReq   map[int]struct{}
}

type reqCfg struct {
	attempts int
	suspend  time.Duration
	preExec  []func()
	postExec []func()
}

type reqOpt func(*reqCfg)

type ReqOptions []reqOpt

func WithReqAttempts(v int) reqOpt {
	return func(rc *reqCfg) {
		rc.attempts = max(v, DEFAULT_FAIL_ATTEMPTS)
	}
}

func WithRetrySuspend(dur time.Duration) reqOpt {
	return func(rc *reqCfg) {
		rc.suspend = max(dur, DEFAULT_RETRY_SUSPEND)
	}
}

func withPreExec(fn func()) reqOpt {
	return func(rc *reqCfg) {
		if fn != nil {
			rc.preExec = append(rc.preExec, fn)
		}
	}
}

func withPostExec(fn func()) reqOpt {
	return func(rc *reqCfg) {
		if fn != nil {
			rc.postExec = append(rc.postExec, fn)
		}
	}
}

func (c *RequestCache) isQryReq(r Request) bool {
	return strings.HasPrefix(r.Executor(), DEFAULT_QRY_REQ_PREFIX)
}

func (c *RequestCache) checkQryInflight() *reqWait {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if len(c.qryInflightReq) >= c.qryInflightCount {
		// 在途查询超限，返回查询等待
		for req := range c.qryInflightReq {
			return c.inflights[req]
		}
	}

	return nil
}

func (c *RequestCache) checkQryFlux() {
	now := time.Now()
	if now.Truncate(time.Second).After(c.lastQry) {
		// 已进入下一周期，查询计数清零
		c.qryLimitCount = 0
	} else if c.qryLimitCount >= c.qryLimit {
		// 查询流控超限，等待下一个周期
		<-time.After(
			now.Add(time.Second).
				Truncate(time.Second).
				Sub(now),
		)
	}
}

func (c *RequestCache) prepareQry(options ReqOptions) ReqOptions {
	if req := c.checkQryInflight(); req != nil {
		req.Wait(-1)
	}

	return append(
		options,
		withPreExec(c.checkQryFlux),
		withPostExec(func() {
			c.lastQry = time.Now().Truncate(time.Second)
			c.qryLimitCount++
		}),
	)
}

func (c *RequestCache) doRequest(r Request, options ...reqOpt) (err error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	var cfg reqCfg

	for _, opt := range options {
		if opt == nil {
			continue
		}

		opt(&cfg)
	}

	for _, pre := range cfg.preExec {
		pre()
	}

	for range cfg.attempts + 1 {
		c.cache = append(c.cache, r)
		seq := len(c.cache)

		if err = r.Execute(seq); err == nil {
			var wait reqWait

			wait.ctx, wait.cancel = context.WithCancel(c.reqRootCtx)
			c.inflights[seq] = &wait

			for _, post := range cfg.postExec {
				post()
			}

			return nil
		}

		slog.Error(
			"request execute failed",
			slog.Any("error", err),
			slog.Int("seq", seq),
			slog.Any("req", r),
		)

		if cfg.suspend > 0 {
			<-time.After(cfg.suspend)
		}
	}

	return err
}

func (c *RequestCache) DoRequestAndWait(
	r Request, options ...reqOpt,
) (*reqWait, error) {
	if c.isQryReq(r) {
		options = c.prepareQry(options)
	}

	var wait *reqWait

	options = append(options, withPostExec(func() {
		wait = c.inflights[len(c.cache)]
	}))

	return wait, c.doRequest(r, options...)
}

func (c *RequestCache) DoRequest(r Request, options ...reqOpt) error {
	if c.isQryReq(r) {
		options = c.prepareQry(options)
	}

	return c.doRequest(r, options...)
}

func (c *RequestCache) Wait(seq int, timeout time.Duration) error {
	if wait := c.getInflight(seq); wait != nil {
		return wait.Wait(timeout)
	}

	return fmt.Errorf("%w: request[%d] completed", ErrInflightNotFound, seq)
}

func (c *RequestCache) SetQryLimit(v int) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.qryLimit = max(v, 1)
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
	delete(c.qryInflightReq, seq)
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
	reqMapper map[reflect.Type]*reflect.Method
	preBuilds map[string]Request
}

func NewRequestFactory[API any](ctx context.Context, api API) *RequestFactory[API] {
	factory := RequestFactory[API]{
		RequestCache: RequestCache{
			qryInflightCount: DEFAULT_QRY_INFLIGHT_REQUESTS,
			// 设置默认每秒1笔查询
			// 避免查询流控检查永远被延迟1s执行
			qryLimit:  1,
			inflights: map[int]*reqWait{},
		},
		Api:       api,
		reqMapper: make(map[reflect.Type]*reflect.Method),
		preBuilds: make(map[string]Request),
	}

	apiType := reflect.TypeFor[API]()

	for fn := range apiType.Methods() {
		if !strings.HasPrefix(fn.Name, "Req") {
			continue
		}

		dataType := fn.Type.In(0)

		factory.reqMapper[dataType] = &fn
	}

	return &factory
}

func (fac *RequestFactory[API]) FilterReqFn(
	filter func(*reflect.Method) bool,
) []*reflect.Method {
	if filter == nil {
		return nil
	}

	results := make([]*reflect.Method, 0, len(fac.reqMapper))

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

func withFactoryRLock[API any, RTN any](
	fac *RequestFactory[API], fn func() (RTN, error),
) (RTN, error) {
	fac.lock.RLock()
	defer fac.lock.RUnlock()

	return fn()
}

func withFactoryLock[API any, RTN any](
	fac *RequestFactory[API], fn func() (RTN, error),
) (RTN, error) {
	fac.lock.Lock()
	defer fac.lock.Unlock()

	return fn()
}

func MakeRequest[
	API any, D thost.ThostData, PTR DataPtr[D],
](factory *RequestFactory[API], v PTR) (Request, error) {
	dataType := reflect.TypeFor[PTR]()

	if req, _ := withFactoryRLock(factory, func() (r Request, e error) {
		r = factory.GetPrebuild(v.Type())
		if r != nil {
			r = r.WithPayload(v)
		}
		return
	}); req != nil {
		return req, nil
	}

	fn, ok := factory.reqMapper[dataType]
	if !ok {
		return nil, fmt.Errorf(
			"%w: no request method for %+v",
			ErrMethodNotFound, dataType,
		)
	}

	req := request[API, D, PTR]{api: factory.Api, payload: v, fnName: fn.Name}

	reflect.ValueOf(&req.handler).Elem().Set(reflect.MakeFunc(
		reflect.TypeOf(req.handler),
		func(args []reflect.Value) (results []reflect.Value) {
			fnImpl := args[0].Method(fn.Index)
			return fnImpl.Call(args[1:])
		},
	))

	return withFactoryLock(factory, func() (Request, error) {
		factory.preBuilds[v.Type()] = &req
		return &req, nil
	})
}
