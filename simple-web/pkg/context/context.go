package context

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpContext struct {
	RspWriter http.ResponseWriter
	Req       *http.Request
}

func (ctx *HttpContext) ReadReqJson(req interface{}) error {
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, req)
}

func (ctx *HttpContext) WriteRspJson(code int, rsp interface{}) error {
	ctx.RspWriter.WriteHeader(code)

	rspJson, err := json.Marshal(rsp)
	if err != nil {
		return err
	}

	_, err = ctx.RspWriter.Write(rspJson)
	return err
}

func (ctx *HttpContext) Ok(rsp interface{}) error {
	return ctx.WriteRspJson(http.StatusOK, rsp)
}

func BuildHttpContext(writer http.ResponseWriter, request *http.Request) *HttpContext {
	return &HttpContext{
		RspWriter: writer,
		Req:       request,
	}
}
