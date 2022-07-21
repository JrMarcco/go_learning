package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Context struct {
	req       *http.Request
	rspWriter http.ResponseWriter

	ctx     context.Context
	handler ControllerHandler

	timeoutFlag bool
	writerMux   *sync.Mutex
}

func NewContext(request *http.Request, responseWriter http.ResponseWriter) *Context {
	return &Context{
		req:       request,
		rspWriter: responseWriter,
		ctx:       request.Context(),
		writerMux: &sync.Mutex{},
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

func (ctx *Context) Err() error {
	return ctx.BaseContext().Err()
}

func (ctx *Context) Value(key any) any {
	return ctx.BaseContext().Value(key)
}

func (ctx *Context) QueryInt(key string, def int) int {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if intValue, err := strconv.Atoi(values[len(values)-1]); err == nil {
			return intValue
		}
	}
	return def
}

func (ctx *Context) QueryString(key string, def string) string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		return values[len(values)-1]
	}
	return def
}

func (ctx Context) QueryArray(key string, def []string) []string {
	params := ctx.QueryAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.URL.Query()
	}
	return map[string][]string{}
}

func (ctx *Context) FormInt(key string, def int) int {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		if intValue, err := strconv.Atoi(values[len(values)-1]); err == nil {
			return intValue
		}
	}
	return def
}

func (ctx *Context) FormString(key string, def string) string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok && len(values) > 0 {
		return values[len(values)-1]
	}
	return def
}

func (ctx *Context) FormArray(key string, def []string) []string {
	params := ctx.FormAll()
	if values, ok := params[key]; ok {
		return values
	}
	return def
}

func (ctx *Context) FormAll() map[string][]string {
	if ctx.req != nil {
		return ctx.req.PostForm
	}
	return map[string][]string{}
}

func (ctx *Context) BindJson(obj any) error {
	if ctx.req != nil {
		body, err := ioutil.ReadAll(ctx.req.Body)
		if err != nil {
			return err
		}
		ctx.req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		if err := json.Unmarshal(body, obj); err != nil {
			return err
		}
	} else {
		return errors.New("ctx.req empty")
	}
	return nil
}

func (ctx *Context) Json(status int, obj any) error {
	if ctx.Timeout() {
		return nil
	}
	ctx.rspWriter.Header().Set("content-Type", "application/json")
	ctx.rspWriter.WriteHeader(status)

	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.rspWriter.WriteHeader(500)
		return err
	}
	if _, err := ctx.rspWriter.Write(byt); err != nil {
		ctx.rspWriter.WriteHeader(500)
		return err
	}
	return nil
}

func (ctx *Context) OkJson(obj any) error {
	return ctx.Json(http.StatusOK, obj)
}

func (ctx *Context) NotFoundJson(obj any) error {
	return ctx.Json(http.StatusNotFound, obj)
}

func (ctx *Context) ErrJson(obj any) error {
	return ctx.Json(http.StatusInternalServerError, obj)
}
