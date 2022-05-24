package server

import (
	"go_learning/simple-web/pkg/route"
	"net/http"
)

type HttpServer interface {
	route.Router
	Start() error
}

type simpleHttpServer struct {
	addr        string
	httpHandler route.RouterHandler
}

func (shs *simpleHttpServer) Route(method string, path string, handleFunc route.HandleFunc) {
	shs.httpHandler.Route(method, path, handleFunc)
}

func (shs *simpleHttpServer) Start() error {
	http.Handle("/", shs.httpHandler)
	return http.ListenAndServe(shs.addr, nil)
}

func newSimpleHttpServer(addr string) HttpServer {
	httpHandler := route.NewTrieTreeHandler()

	return &simpleHttpServer{
		addr:        addr,
		httpHandler: httpHandler,
	}
}

func DefaultHttpServer() HttpServer {
	return newSimpleHttpServer(":8080")
}
