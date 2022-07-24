package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct {
	req       *http.Request
	rspWriter http.ResponseWriter

	ctx          context.Context
	handlerIndex int
	handlers     HandlerChain

	timeoutFlag bool
	writerMux   *sync.Mutex

	params map[string]string
}

func NewContext(request *http.Request, responseWriter http.ResponseWriter) *Context {
	return &Context{
		req:          request,
		rspWriter:    responseWriter,
		ctx:          request.Context(),
		writerMux:    &sync.Mutex{},
		handlerIndex: -1,
	}
}

func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

func (ctx *Context) Request() *http.Request {
	return ctx.req
}

func (ctx *Context) ResponseWriter() http.ResponseWriter {
	return ctx.rspWriter
}

func (ctx *Context) SetTimeout() {
	ctx.timeoutFlag = true
}

func (ctx *Context) SetHandlers(handlers []HandlerFunc) {
	ctx.handlers = handlers
}

func (ctx *Context) SetParams(params map[string]string) {
	ctx.params = params
}

func (ctx *Context) Timeout() bool {
	return ctx.timeoutFlag
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.req.Context()
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.BaseContext().Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Next() error {
	ctx.handlerIndex++
	if ctx.handlerIndex < len(ctx.handlers) {
		if err := ctx.handlers[ctx.handlerIndex](ctx); err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}
