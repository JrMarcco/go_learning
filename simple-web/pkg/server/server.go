package server

import (
	"go_learning/simple-web/pkg/context"
	"go_learning/simple-web/pkg/filter"
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
	rootFilter  filter.HttpFiler
}

func (shs *simpleHttpServer) Route(method string, path string, handleFunc route.HandleFunc) {
	shs.httpHandler.Route(method, path, handleFunc)
}

func (shs *simpleHttpServer) Start() error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ctx := context.BuildHttpContext(writer, request)
		shs.rootFilter(ctx)
	})
	return http.ListenAndServe(shs.addr, nil)
}

func newSimpleHttpServer(addr string, httpFilterBuilders []filter.HttpFilterBuilder) HttpServer {
	httpHandler := route.NewTrieTreeHandler()
	var rootFilter filter.HttpFiler = func(ctx *context.HttpContext) {
		httpHandler.ServeHTTP(ctx.RspWriter, ctx.Req)
	}

	if len(httpFilterBuilders) > 0 {
		for i := len(httpFilterBuilders); i >= 0; i-- {
			builder := httpFilterBuilders[i]
			rootFilter = builder(rootFilter)
		}
	}

	return &simpleHttpServer{
		addr:        addr,
		httpHandler: httpHandler,
		rootFilter:  rootFilter,
	}
}

func DefaultHttpServer(httpFilterBuilders ...filter.HttpFilterBuilder) HttpServer {
	return newSimpleHttpServer(":8080", httpFilterBuilders)
}
