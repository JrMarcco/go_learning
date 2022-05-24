package ctx

import (
	"encoding/json"
	"io"
	"net/http"
)

type HttpContext struct {
	RspWriter http.ResponseWriter
	Req       *http.Request
}

func (ctx *HttpContext) ReadReqJson(req any) error {
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, req)
}

func (ctx *HttpContext) WriteRspJson(code int, rsp any) error {
	ctx.RspWriter.WriteHeader(code)

	if rsp != nil {
		rspJson, err := json.Marshal(rsp)
		if err != nil {
			return err
		}

		_, err = ctx.RspWriter.Write(rspJson)
		return err
	}
	return nil
}

func (ctx *HttpContext) Ok(rsp any) error {
	return ctx.WriteRspJson(http.StatusOK, rsp)
}

func (ctx *HttpContext) Bad() error {
	return ctx.WriteRspJson(http.StatusBadRequest, nil)
}

func (ctx *HttpContext) NotFound() error {
	return ctx.WriteRspJson(http.StatusNotFound, nil)
}

func BuildHttpContext(writer http.ResponseWriter, request *http.Request) *HttpContext {
	return &HttpContext{
		RspWriter: writer,
		Req:       request,
	}
}
