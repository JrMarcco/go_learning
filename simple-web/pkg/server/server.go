package server

import (
	"net/http"
)

type HttpServer interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start() error
}

type simpleHttpServer struct {
	addr string
}

func (shs *simpleHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc("/", handlerFunc)
}

func (shs *simpleHttpServer) Start() error {
	return http.ListenAndServe(shs.addr, nil)
}

func NewSimpleHttpServer(addr string) HttpServer {
	if addr == "" {
		addr = ":8080"
	}

	return &simpleHttpServer{
		addr: addr,
	}
}
