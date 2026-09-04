package state

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"reflect"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/frozenpine/ctp4go"
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

type reqFactoryCtxKey struct{}

var (
	ReqSeqBaseKey reqFactoryCtxKey = struct{}{}
)

type Request interface {
	fmt.Stringer

	Type() string
	Executor() string

	WithPayload(any) Request
	Execute(int) error
}

func CastData[
	REQ ctp4go.DataConstraint, P_REQ DataPtr[REQ],
	RSP ctp4go.DataConstraint, P_RSP DataPtr[RSP],
](r Request) (P_REQ, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}

	v, ok := r.(*request[REQ, P_REQ])
	if !ok {
		return nil, errors.New("unsupported request")
	}

	return v.payload, nil
}

type request[
	REQ ctp4go.DataConstraint, P_REQ DataPtr[REQ],
] struct {
	fnName  string
	handler func(P_REQ, int) int

	payload P_REQ
}

func (r *request[REQ, P_REQ]) Type() string {
	return r.payload.Type()
}

func (r *request[REQ, P_REQ]) Executor() string { return r.fnName }

func (r *request[REQ, P_REQ]) WithPayload(v any) Request {
	if p, ok := v.(P_REQ); ok {
		newReq := *r
		newReq.payload = p
		return &newReq
	}

	slog.Error(
		"payload mismatch with request",
		slog.Any("req_type", reflect.TypeFor[P_REQ]()),
		slog.Any("payload_type", reflect.TypeOf(v)),
	)

	return nil
}

func (r *request[REQ, P_REQ]) Execute(seq int) error {
	slog.Debug(
		"start to execute request",
		slog.Int("seq", seq),
		slog.String("executor", r.fnName),
		slog.Any("payload", r.payload),
	)

	return ctp4go.Rtn(r.handler(r.payload, seq)).Error()
}

func (r *request[REQ, P_REQ]) String() string {
	builder := ctp4go.GetStringBuilder()

	fmt.Fprintf(
		builder, "Request{Executor=%s, Payload=%+v}",
		r.fnName, r.payload,
	)

	return builder.String()
}

type reqWait struct {
	ctx          context.Context
	cancel       context.CancelFunc
	err          error
	completeExec []RspCallback
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
		return errors.Join(waitCtx.Err(), w.err)
	}
}

type RequestCache struct {
	lock       sync.RWMutex
	reqRootCtx context.Context

	seqBase int
	cache   []Request
	// inflights map[int]*reqWait
	inflights sync.Map

	lastQry          time.Time
	qryInflightCount int
	qryLimit         int
	qryLimitCount    int
	// qryInflightReq   map[int]struct{}
	qryInflightReq sync.Map
}

type reqCfg struct {
	attempts     int
	suspend      time.Duration
	preExec      []func()
	postExec     []func()
	completeExec []RspCallback
}

type reqOpt func(*reqCfg)

type ReqOptions []reqOpt

type RspCallback func(ctp4go.DataConstraint, error, int, bool)

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

func WithPreExec(fn func()) reqOpt {
	return func(rc *reqCfg) {
		if fn == nil {
			slog.Warn("pre execution is nil")
			return
		}

		rc.preExec = append(rc.preExec, fn)
	}
}

func WithPostExec(fn func()) reqOpt {
	return func(rc *reqCfg) {
		if fn == nil {
			slog.Warn("post execution is nil")
			return
		}

		rc.postExec = append(rc.postExec, fn)
	}
}

func WithCompleteExec(fn RspCallback) reqOpt {
	return func(rc *reqCfg) {
		if fn == nil {
			slog.Warn("complete execution is nil")
			return
		}

		rc.completeExec = append(rc.completeExec, fn)
	}
}

// func (c *RequestCache) Lock() { c.lock.Lock() }

// func (c *RequestCache) Unlock() { c.lock.Unlock() }

// func (c *RequestCache) RLock() { c.lock.RLock() }

// func (c *RequestCache) RUnlock() { c.lock.RUnlock() }

func (c *RequestCache) isQryReq(r Request) bool {
	return strings.HasPrefix(r.Executor(), DEFAULT_QRY_REQ_PREFIX)
}

func (c *RequestCache) checkQryInflight() *reqWait {
	// c.lock.RLock()
	// defer c.lock.RUnlock()

	// if len(c.qryInflightReq) >= c.qryInflightCount {
	// 	// 在途查询超限，返回查询等待
	// 	for req := range c.qryInflightReq {
	// 		return c.inflights[req]
	// 	}
	// }

	var (
		qryInflight = 0
		qry         *reqWait
	)

	c.qryInflightReq.Range(func(key, value any) bool {
		qryInflight++

		if qryInflight >= c.qryInflightCount {
			qry = value.(*reqWait)
			return false
		}

		return true
	})

	return qry
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
		WithPreExec(c.checkQryFlux),
		WithPostExec(func() {
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

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"request options applied",
		slog.Any("request", r),
		slog.Int("attempts", cfg.attempts),
		slog.Duration("suspend", cfg.suspend),
		slog.Int("pre", len(cfg.preExec)),
		slog.Int("post", len(cfg.postExec)),
		slog.Int("complete", len(cfg.completeExec)),
	)

	for _, pre := range cfg.preExec {
		pre()
	}

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"request pre execution applied",
		slog.Any("request", r),
	)

	for idx := range cfg.attempts + 1 {
		c.cache = append(c.cache, r)
		seq := c.getSeq()

		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"attemp to execute request",
			slog.Int("attempts", idx+1),
			slog.Any("request", r),
			slog.Int("seq", seq),
		)

		if err = r.Execute(seq); err == nil {
			wait := reqWait{
				completeExec: cfg.completeExec,
			}

			wait.ctx, wait.cancel = context.WithCancel(c.reqRootCtx)
			// c.inflights[seq] = &wait
			c.inflights.Store(seq, &wait)

			for _, post := range cfg.postExec {
				post()
			}

			slog.Debug(
				"request executed",
				slog.Any("req", r),
				slog.Int("seq", seq),
			)

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

	options = append(options, WithPostExec(func() {
		// 提取在途请求的wait对象
		// wait = c.inflights[c.getSeq()]
		wait = c.getInflight(c.getSeq())
	}))

	return wait, c.doRequest(r, options...)
}

func (c *RequestCache) getSeq() int {
	return len(c.cache) + c.seqBase
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
	// c.lock.RLock()
	// defer c.lock.RUnlock()

	// return c.inflights[seq]
	if v, ok := c.inflights.Load(seq); ok {
		return v.(*reqWait)
	}
	return nil
}

func (c *RequestCache) WithResponse(
	data ctp4go.DataConstraint, rsp error,
	seq int, last bool,
) {
	reqWait := c.getInflight(seq)
	if reqWait == nil {
		slog.Warn(
			"response without inflight request",
			slog.Int("seq", seq),
			slog.Bool("last", last),
			slog.Any("data", data),
			slog.Any("rsp", rsp),
		)
		return
	}

	defer func() {
		reqWait.cancel()

		for _, cb := range reqWait.completeExec {
			cb(data, rsp, seq, last)
		}
	}()

	if rsp != nil && reqWait.err == nil {
		reqWait.err = rsp
	}

	if last {
		// c.lock.Lock()
		// defer c.lock.Unlock()

		// delete(c.inflights, seq)
		c.inflights.Delete(seq)
		// delete(c.qryInflightReq, seq)
		c.qryInflightReq.Delete(seq)
	}
}

func (c *RequestCache) Reset() {
	c.lock.Lock()
	defer c.lock.Unlock()

	// c.inflights = make(map[int]*reqWait)
	c.inflights.Clear()
	c.qryInflightReq.Clear()
	c.cache = c.cache[:0]
}

type RCache interface {
	// RLock()
	// RUnlock()
	// Lock()
	// Unlock()

	DoRequestAndWait(Request, ...reqOpt) (*reqWait, error)
	DoRequest(Request, ...reqOpt) error
	Wait(seq int, timeout time.Duration) error
	WithResponse(ctp4go.DataConstraint, error, int, bool)
	Reset()
}

type RFactory interface {
	RCache

	GetReqMethod(string) *ReqMethod
	GetPrebuild(string) Request

	SetPrebuilds(string, Request)
}

var (
	_ RFactory = (*RequestFactory[int])(nil)
)

type ReqMethod struct {
	name string
	reflect.Value
}

func (m ReqMethod) MethodName() string { return m.name }

type RequestFactory[API any] struct {
	RequestCache

	Api        API
	reqMethods map[string]*ReqMethod
	// preBuilds  map[string]Request
	preBuilds sync.Map
}

func NewRequestFactory[API any](
	ctx context.Context, api API,
) *RequestFactory[API] {
	var seqBase int

	if ctx == nil {
		ctx = context.Background()
	} else {
		if v, ok := ctx.Value(ReqSeqBaseKey).(int); ok {
			seqBase = v
		}
	}

	factory := RequestFactory[API]{
		RequestCache: RequestCache{
			reqRootCtx:       ctx,
			seqBase:          seqBase,
			qryInflightCount: DEFAULT_QRY_INFLIGHT_REQUESTS,
			// 设置默认每秒1笔查询
			// 避免查询流控检查永远被延迟1s执行
			qryLimit: 1,
			// inflights: map[int]*reqWait{},
		},
		Api:        api,
		reqMethods: make(map[string]*ReqMethod),
		// preBuilds:  make(map[string]Request),
	}

	apiValue := reflect.ValueOf(api)

	for fn, ins := range apiValue.Methods() {
		if !strings.HasPrefix(fn.Name, "Req") {
			continue
		}

		args := slices.Collect(fn.Type.Ins())
		rtn := slices.Collect(fn.Type.Outs())

		if len(args) != 3 || (args[1].Kind() != reflect.Pointer ||
			args[1].Elem().Kind() != reflect.Struct) ||
			args[2].Kind() != reflect.Int {
			slog.Error(
				"invalid request method args",
				slog.String("executor", fn.Name),
				slog.Any("args", args),
			)
			continue
		}

		if len(rtn) != 1 || rtn[0].Kind() != reflect.Int {
			slog.Error(
				"invalid request method rtn",
				slog.String("executor", fn.Name),
				slog.Any("rtn", rtn),
			)
			continue
		}

		slog.Info(
			"request method found",
			slog.String("executor", fn.Name),
			slog.Any("args", args),
			slog.Any("rtn", rtn[0]),
		)

		factory.reqMethods[args[1].Elem().Name()] = &ReqMethod{
			name:  fn.Name,
			Value: ins,
		}
	}

	return &factory
}

func (fac *RequestFactory[API]) GetReqMethod(
	dataType string,
) *ReqMethod {
	return fac.reqMethods[dataType]
}

func (fac *RequestFactory[API]) GetPrebuild(name string) Request {
	// fac.lock.RLock()
	// defer fac.lock.RUnlock()

	// return fac.preBuilds[name]
	if v, ok := fac.preBuilds.Load(name); ok {
		return v.(Request)
	}

	return nil
}

func (fac *RequestFactory[API]) SetPrebuilds(typName string, r Request) {
	// fac.preBuilds[typName] = r
	fac.preBuilds.CompareAndSwap(typName, nil, r)
}

// func withFactoryRLock[RTN any](
// 	fac RFactory, fn func() (RTN, error),
// ) (RTN, error) {
// 	fac.RLock()
// 	defer fac.RUnlock()

// 	return fn()
// }

// func withFactoryLock[RTN any](
// 	fac RFactory, fn func() (RTN, error),
// ) (RTN, error) {
// 	fac.Lock()
// 	defer fac.Unlock()

// 	return fn()
// }

func MakeRequest[
	REQ ctp4go.DataConstraint, P_REQ DataPtr[REQ],
](
	factory RFactory, payload P_REQ,
) (Request, error) {
	// if req, _ := withFactoryRLock(factory, func() (r Request, e error) {
	// 	r = factory.GetPrebuild(payload.Type())
	// 	if r != nil {
	// 		r = r.WithPayload(payload)
	// 	}
	// 	return
	// }); req != nil {
	// 	return req, nil
	// }
	if r := factory.GetPrebuild(payload.Type()); r != nil {
		r = r.WithPayload(payload)
		return r, nil
	}

	fn := factory.GetReqMethod(payload.Type())
	if fn == nil || !fn.IsValid() {
		return nil, fmt.Errorf(
			"%w: no request method for %+v",
			ErrMethodNotFound, payload.Type(),
		)
	}

	req := request[REQ, P_REQ]{
		fnName:  fn.MethodName(),
		payload: payload,
	}

	var ok bool
	req.handler, ok = fn.Interface().(func(P_REQ, int) int)
	if !ok {
		return nil, fmt.Errorf(
			"%w: assert request method[%s] failed",
			ErrMethodNotFound, fn.Type().Name(),
		)
	}

	// return withFactoryLock(factory, func() (Request, error) {
	// 	factory.SetPrebuilds(payload.Type(), &req)
	// 	return &req, nil
	// })
	factory.SetPrebuilds(payload.Type(), &req)
	return &req, nil
}
