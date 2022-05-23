package server

import (
	"go_learning/simple-web/pkg/context"
	"net/http"
)

type HttpServer interface {
	Route(pattern string, handlerFunc func(ctx *context.HttpContext))
	Start() error
}

type simpleHttpServer struct {
	addr string
}

func (shs *simpleHttpServer) Route(pattern string, handlerFunc func(ctx *context.HttpContext)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		handlerFunc(context.BuildHttpContext(writer, request))
	})
}

func (shs *simpleHttpServer) Start() error {
	return http.ListenAndServe(shs.addr, nil)
}

func newSimpleHttpServer(addr string) HttpServer {
	return &simpleHttpServer{
		addr: addr,
	}
}

func DefaultHttpServer() HttpServer {
	return newSimpleHttpServer(":8080")
}
