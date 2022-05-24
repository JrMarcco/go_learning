package route

import "go_learning/simple-web/pkg/ctx"

type HandleFunc func(ctx *ctx.HttpContext)

type Router interface {
	Route(method string, path string, handleFunc HandleFunc)
}
