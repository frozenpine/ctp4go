package state

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidChain    = errors.New("invalid chain")
	ErrInvalidChainReq = errors.New("invalid chain request")
	ErrReqSpanFailed   = errors.New("request span failed")
	ErrReqSpanBreaked  = errors.New("request span breaked")
	ErrRspSpanFailed   = errors.New("response span failed")
	ErrRspSpanTimeout  = errors.New("response span timeout")
	ErrSpanAlreadyDone = errors.New("span already done")
)

//go:generate stringer -type spanState -linecomment
type spanState uint8

const (
	SpanCreated       spanState = iota // 已创建
	SpanPrepared                       // 已就绪
	SpanPrepareFailed                  // 准备失败
	SpanReqSuccess                     // 已执行
	SpanReqFailed                      // 执行失败
	SpanRspArrived                     // 回执到达
	SpanRspError                       // 回执错误
	SpanRspFinished                    // 已完成
	SpanBreaked                        // 中断
	SpanTimeouted                      // 超时
	SpanBreakContinue                  // 中断继续
)

var terminatedStates = []spanState{
	SpanPrepareFailed,
	SpanReqFailed,
	SpanRspError,
	SpanBreaked,
	SpanRspFinished,
	SpanBreakContinue,
}

func (curr *spanState) Migrate(next spanState) error {
	switch next {
	case SpanPrepared, SpanPrepareFailed:
		if *curr != SpanCreated {
			goto ERR
		}
	case SpanReqSuccess:
		switch *curr {
		case SpanPrepared:
		case SpanRspArrived, SpanRspError, SpanRspFinished,
			SpanBreaked, SpanBreakContinue:
			// 增加已有回报状态判断
			// 以避免回执比请求成功状态更早到达
			return nil
		default:
			goto ERR
		}
	case SpanReqFailed:
		switch *curr {
		case SpanCreated, SpanPrepared:
		default:
			goto ERR
		}
	case SpanRspArrived, SpanRspError, SpanRspFinished,
		SpanBreakContinue, SpanBreaked:
		switch *curr {
		// 增加已就绪前置状态，以避免回执比请求成功状态更早到达
		case SpanPrepared, SpanReqSuccess, SpanRspArrived:
		default:
			goto ERR
		}
	default:
		goto ERR
	}

	*curr = next

	return nil

ERR:
	return fmt.Errorf(
		"%w: can not migrate to %s from %s",
		ErrInvalidChain, next.String(), curr.String(),
	)
}

type PreSpan func() error
type SpanExec func(*SpanInfo) error
type PostSpan func(error) error

type SpanInfo struct {
	ReqID  string `json:"reqID"`
	SpanID string `json:"spanID"`
	lock   sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	chain *Executor
	prev  *SpanInfo
	next  *SpanInfo

	prepares []PreSpan
	exec     SpanExec
	posts    []PostSpan
	execErr  error

	state *FlagResponsor[spanState, any]
}

func (e *SpanInfo) Execute() (err error) {
	e.chain.lock.Lock()
	if !e.chain.concurrent {
		slog.Debug(
			"rotate executor current span",
			slog.String("req_id", e.chain.ReqID),
			slog.String("span_id", e.SpanID),
		)
		e.chain.currentSpan = e
	}
	e.chain.lock.Unlock()

	if err := e.Prepare(); err != nil {
		return err
	}

	slog.Debug(
		"executing span",
		slog.String("req_id", e.ReqID),
		slog.String("span_id", e.SpanID),
		slog.String("state", e.state.String()),
	)

	defer func() {
		if err := recover(); err != nil {
			slog.Error(
				"span execution recovered",
				slog.Any("error", err),
				slog.String("req_id", e.ReqID),
				slog.String("span_id", e.SpanID),
			)
		}
	}()

	e.lock.Lock()
	err = e.exec(e)
	e.lock.Unlock()

	if err != nil {
		slog.Error(
			"span executed with error",
			slog.Any("error", err),
			slog.String("req_id", e.ReqID),
			slog.String("span_id", e.SpanID),
			slog.String("state", e.state.String()),
		)
	}

	if err = e.PostExec(err); err != nil {
		e.execErr = err
		return errors.Join(ErrReqSpanFailed, err, e.state.SetFlag(SpanReqFailed))
	}

	if err = e.state.SetFlag(SpanReqSuccess); err != nil {
		return errors.Join(ErrReqSpanFailed, err)
	} else {
		slog.Debug(
			"span executed",
			slog.String("req_id", e.ReqID),
			slog.String("span_id", e.SpanID),
		)
	}

	return nil
}

func (e *SpanInfo) Prepare() (err error) {
	slog.Debug(
		"preparing execution span",
		slog.String("req_id", e.ReqID),
		slog.String("span_id", e.SpanID),
		slog.String("state", e.state.String()),
	)

	defer func() {
		if err := recover(); err != nil {
			slog.Error(
				"span prepare recovered with error",
				slog.String("req_id", e.ReqID),
				slog.String("span_id", e.SpanID),
				slog.Any("error", err),
			)
		}
	}()

	for _, p := range e.prepares {
		if p == nil {
			continue
		}

		if err = p(); err != nil {
			slog.Debug(
				"execution span prepare failed",
				slog.Any("error", err),
				slog.String("req_id", e.ReqID),
				slog.String("span_id", e.SpanID),
				slog.String("state", e.state.String()),
			)
			return errors.Join(
				ErrReqSpanFailed, err, e.state.SetFlag(SpanPrepareFailed),
			)
		}
	}

	if err := e.state.SetFlag(SpanPrepared); err != nil {
		slog.Error(
			"execution span prepared with state rotate fail",
			slog.Any("error", err),
			slog.String("req_id", e.ReqID),
			slog.String("span_id", e.SpanID),
			slog.String("state", e.state.String()),
		)
		return errors.Join(ErrReqSpanFailed, err)
	}

	slog.Debug(
		"execution span prepared",
		slog.String("req_id", e.ReqID),
		slog.String("span_id", e.SpanID),
		slog.String("state", e.state.String()),
	)

	return nil
}

func (e *SpanInfo) PostExec(execErr error) error {
	if len(e.posts) < 1 {
		return execErr
	}

	slog.Debug(
		"post execution span",
		slog.String("req_id", e.ReqID),
		slog.String("span_id", e.SpanID),
		slog.String("state", e.state.String()),
	)

	defer func() {
		if err := recover(); err != nil {
			slog.Error(
				"span post execution recovered with error",
				slog.String("req_id", e.ReqID),
				slog.String("span_id", e.SpanID),
				slog.Any("error", err),
			)
		}
	}()

	postErrs := []error{}

	for _, p := range e.posts {
		if p == nil {
			continue
		}

		if err := p(execErr); err != nil {
			postErrs = append(postErrs, err)
		}
	}

	if err := errors.Join(postErrs...); err == nil {
		if execErr != nil {
			slog.Info(
				"execution span post recoverd exec error",
				slog.Any("exec_error", execErr),
				slog.String("req_id", e.ReqID),
				slog.String("span_id", e.SpanID),
				slog.String("state", e.state.String()),
			)
		}
		return nil
	} else {
		slog.Error(
			"execution span post handled with err",
			slog.Any("exec_error", execErr),
			slog.Any("error", err),
			slog.String("req_id", e.ReqID),
			slog.String("span_id", e.SpanID),
			slog.String("state", e.state.String()),
		)
		return err
	}
}

func (e *SpanInfo) linkNext(n *SpanInfo) {
	if n == nil {
		return
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.next != nil {
		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"link span break chain",
			slog.String("src_span", e.SpanID),
			slog.String("dst_name", e.next.SpanID),
		)

		e.next.lock.Lock()
		defer e.next.lock.Unlock()

		e.next.prev, n.prev = n, e
		e.next, n.next = n, e.next

		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"link span next inject",
			slog.String("src_span", n.prev.SpanID),
			slog.String("mid_span", n.SpanID),
			slog.String("dst_span", n.next.SpanID),
		)
	} else {
		e.next = n
		n.prev = e
		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"new span linked",
			slog.String("src_span", n.prev.SpanID),
			slog.String("dst_span", n.SpanID),
		)
	}
}

func (e *SpanInfo) InjectSpan(req SpanExec, options ...spanOpt) (*SpanInfo, error) {
	e.chain.lock.Lock()
	defer e.chain.lock.Unlock()

	if err := e.chain.AppendSpan(
		WithExecSpan(req, options...),
	); err != nil {
		return nil, err
	}

	newSpan := e.chain.spanList[len(e.chain.spanList)-1]

	e.linkNext(newSpan)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"new span injected",
		slog.String("span", newSpan.SpanID),
	)
	return newSpan, nil
}

func (e *SpanInfo) Prev() *SpanInfo {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.prev
}

func (e *SpanInfo) Next() *SpanInfo {
	e.lock.RLock()
	defer e.lock.RUnlock()
	return e.next
}

func (e *SpanInfo) Payload() any {
	return e.state.GetPayload()
}

func (e *SpanInfo) SetPayload(v any) {
	e.state.SetPayload(v)
}

func (e *SpanInfo) SetFlag(v spanState, payload ...any) error {
	return e.state.SetFlag(v, payload...)
}

func (e *SpanInfo) GetFlag() spanState {
	return e.state.GetFlag()
}

func (e *SpanInfo) Done() <-chan struct{} {
	return e.ctx.Done()
}

func (e *SpanInfo) Break() error {
	select {
	case <-e.Done():
		return fmt.Errorf("%w: %s", ErrSpanAlreadyDone, e.SpanID)
	default:
		return e.SetFlag(SpanBreaked)
	}
}

func GetSpanPayload[T any](span *SpanInfo) (T, error) {
	var v T
	if span == nil {
		return v, fmt.Errorf("%w: span is nil", ErrInvalidChainReq)
	}

	var ok bool

	v, ok = span.Payload().(T)
	if !ok {
		return v, fmt.Errorf("%w state payload mismatch", ErrInvalidChainReq)
	}

	return v, nil
}

type execOpt func(*Executor) error
type ExecOptions []execOpt

type spanOpt func(*SpanInfo) error
type SpanOptions []spanOpt

func WithSpanPrep(fn PreSpan) spanOpt {
	return func(si *SpanInfo) error {
		if fn == nil {
			return fmt.Errorf(
				"%w: prep func is nil", ErrInvalidChain,
			)
		}

		si.prepares = append(si.prepares, fn)
		return nil
	}
}

func WithSpanPost(fn PostSpan) spanOpt {
	return func(si *SpanInfo) error {
		if fn == nil {
			return fmt.Errorf(
				"%w: post func is nil", ErrInvalidChain,
			)
		}

		si.posts = append(si.posts, fn)
		return nil
	}
}

func WithExecSpan(
	req SpanExec, options ...spanOpt,
) execOpt {
	return func(dc *Executor) error {
		if req == nil || dc == nil {
			return fmt.Errorf("%w: name or requst empty", ErrInvalidChain)
		}

		spanID, err := uuid.NewV7()
		if err != nil {
			return fmt.Errorf(
				"%w: gen span id failed %+v", ErrInvalidChain, err,
			)
		}

		span := &SpanInfo{
			ReqID:  dc.ReqID,
			SpanID: spanID.String(),
			chain:  dc,
			exec:   req,
		}

		for _, opt := range options {
			if opt == nil {
				continue
			}

			if err := opt(span); err != nil {
				return err
			}
		}

		span.ctx, span.cancel = context.WithCancel(dc.execCtx)
		span.state = NewPayloadFlagResponsor[spanState, any](span.SpanID)

		// 注入全部终止状态回调
		for _, s := range terminatedStates {
			hook, err := NewHandler(func() error {
				slog.Log(
					context.Background(), slog.LevelDebug-1,
					"terminate state triggered for span",
					slog.String("span_id", span.SpanID),
					slog.String("state", s.String()),
				)

				switch s {
				case SpanRspFinished, SpanBreakContinue:
					if span.next != nil {
						go span.next.Execute()
					}
				}

				span.cancel()
				slog.Log(
					context.Background(), slog.LevelDebug-1,
					"span execution cancelled",
					slog.String("span_id", span.SpanID),
				)
				return nil
			})
			if err != nil {
				return errors.Join(ErrInvalidChain, err)
			}

			if err = span.state.AddHandler(s, hook); err != nil {
				return errors.Join(ErrInvalidChain, err)
			}
		}

		dc.spanList = append(dc.spanList, span)
		dc.cache.Store(spanID, span)

		return nil
	}
}

func WithExecSpanPayload[T any](
	req SpanExec, payload T, options ...spanOpt,
) execOpt {
	return WithExecSpan(req, append(options, func(si *SpanInfo) error {
		si.state.SetPayload(payload)
		return nil
	})...)
}

func WithDelaySpan(
	delay time.Duration, req SpanExec, options ...spanOpt,
) execOpt {
	return WithExecSpan(req, append(options, WithSpanPrep(
		func() error {
			if delay <= 0 {
				return fmt.Errorf(
					"%w: delay time invalid %+v", ErrInvalidChain, delay,
				)
			}

			<-time.After(delay)
			return nil
		},
	))...)
}

type Executor struct {
	Name  string `json:"name"`
	ReqID string `json:"reqID"`

	execCtx    context.Context
	execCancel context.CancelFunc

	lock        sync.RWMutex
	concurrent  bool
	currentSpan *SpanInfo
	spanList    []*SpanInfo
	cache       sync.Map
}

func NewExecutor(ctx context.Context, name string) (*Executor, error) {
	reqID, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf(
			"%w: gen request id failed %+v", ErrInvalidChain, err,
		)
	}

	if ctx == nil {
		ctx = context.Background()
	}

	exec := &Executor{
		Name:  name,
		ReqID: reqID.String(),
	}

	exec.execCtx, exec.execCancel = context.WithCancel(ctx)

	return exec, nil
}

func (chain *Executor) Wait(
	id string, timeout time.Duration,
) error {
	spanId, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf(
			"%w: parse span id failed",
			errors.Join(ErrInvalidChainReq, err),
		)
	}
	chain.lock.RLock()
	e, exist := chain.cache.Load(spanId)
	chain.lock.RUnlock()
	if !exist {
		return fmt.Errorf("%w: name %s", ErrInvalidChainReq, id)
	}

	exec := e.(*SpanInfo)

	waitCtx := context.Background()

	if timeout > 0 {
		var cancel context.CancelFunc

		waitCtx, cancel = context.WithTimeout(waitCtx, timeout)
		defer cancel()
	}

	select {
	case <-exec.Done():
		return nil
	case <-waitCtx.Done():
		return fmt.Errorf(
			"%w: wait %s for %+v",
			ErrRspSpanTimeout, id, timeout,
		)
	}
}

func (chain *Executor) WaitAll(timeout time.Duration) (err error) {
	waitCtx := context.Background()

	if timeout > 0 {
		var cancel context.CancelFunc

		waitCtx, cancel = context.WithTimeout(waitCtx, timeout)
		defer cancel()
	}

	chain.lock.RLock()
	beginCount := len(chain.spanList)
	chain.lock.RUnlock()

	for {
		chain.cache.Range(func(key, value any) bool {
			span := value.(*SpanInfo)

			slog.Log(
				context.Background(), slog.LevelDebug-1,
				"waiting for span",
				slog.String("span", span.SpanID),
			)

			select {
			case <-span.Done():
				switch span.state.GetFlag() {
				case SpanPrepareFailed, SpanReqFailed:
					err = errors.Join(ErrReqSpanFailed, span.execErr)
				case SpanRspError:
					err = errors.Join(ErrRspSpanFailed, span.execErr)
				case SpanBreaked:
					err = errors.Join(ErrReqSpanBreaked, span.execErr)
				case SpanRspFinished:
					return true
				}
			case <-waitCtx.Done():
				err = ErrRspSpanTimeout
				return false
			}

			return false
		})

		chain.lock.RLock()
		endCount := len(chain.spanList)
		chain.lock.RUnlock()
		if endCount == beginCount {
			break
		}

		slog.Info(
			"new span injected when waiting",
			slog.Int("before_waiting", beginCount),
			slog.Int("current", endCount),
		)
		beginCount = endCount
	}

	return
}

func (chain *Executor) AllSpans() []*SpanInfo {
	chain.lock.RLock()
	defer chain.lock.RUnlock()
	return chain.spanList
}

func (chain *Executor) AppendSpan(spans ...execOpt) error {
	for _, spanOpt := range spans {
		if spanOpt == nil {
			continue
		}

		if err := spanOpt(chain); err != nil {
			return err
		}
	}

	return nil
}

func (chain *Executor) Reset(spans ...execOpt) error {
	chain.lock.Lock()
	defer chain.lock.Unlock()

	chain.spanList = chain.spanList[:0]
	chain.cache.Clear()

	return chain.AppendSpan(spans...)
}

// ChainExecute 顺序异步执行
func (chain *Executor) ChainExecute(spans ...execOpt) error {
	chain.lock.Lock()
	defer chain.lock.Unlock()

	if err := chain.AppendSpan(spans...); err != nil {
		return err
	}
	if len(chain.spanList) < 1 {
		return ErrInvalidChainReq
	}

	chain.concurrent = false

	for idx, e := range chain.spanList {
		if e.state == nil {
			return fmt.Errorf(
				"%w: %s[%d] execution is nil",
				ErrInvalidChainReq, e.SpanID, idx,
			)
		}

		if idx < 1 {
			continue
		}

		prev := chain.spanList[idx-1]
		if prev.state == nil {
			return fmt.Errorf(
				"%w: %s[%d]'s pre execution is nil %s[%d]",
				ErrInvalidChainReq, e.SpanID, idx, prev.SpanID, idx-1,
			)
		}
		prev.linkNext(e)

	}

	go chain.spanList[0].Execute()

	return nil
}

// ChainExecuteAndWait 顺序同步执行
func (chain *Executor) ChainExecuteAndWait(
	timeout time.Duration, executions ...execOpt,
) error {
	if err := chain.ChainExecute(executions...); err != nil {
		return err
	}

	return chain.WaitAll(timeout)
}

// ConcurrentExecute 并发异步执行
func (chain *Executor) ConcurrentExecute(executions ...execOpt) error {
	chain.lock.Lock()
	defer chain.lock.Unlock()

	if err := chain.AppendSpan(executions...); err != nil {
		return err
	}
	if len(chain.spanList) < 1 {
		return ErrInvalidChainReq
	}

	chain.concurrent = true

	for idx, e := range chain.spanList {
		if e.state == nil {
			return fmt.Errorf(
				"%w: %s[%d] execution is nil",
				ErrInvalidChainReq, e.SpanID, idx,
			)
		}

		go e.Execute()
	}

	return nil
}

// ConcurentExecuteAndWait 并发同步执行
func (chain *Executor) ConcurentExecuteAndWait(
	timeout time.Duration, executions ...execOpt,
) error {
	if err := chain.ConcurrentExecute(executions...); err != nil {
		return err
	}

	return chain.WaitAll(timeout)
}

func (chain *Executor) CurrentSpan() *SpanInfo {
	chain.lock.RLock()
	defer chain.lock.RUnlock()

	return chain.currentSpan
}
