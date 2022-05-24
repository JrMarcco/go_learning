package route

import "go_learning/simple-web/pkg/hctx"

type HandleFunc func(ctx *hctx.HttpContext)

type Router interface {
	Route(method string, path string, handleFunc HandleFunc)
}
