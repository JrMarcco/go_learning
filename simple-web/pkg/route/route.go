package route

import "go_learning/simple-web/pkg/context"

type HandleFunc func(ctx *context.HttpContext)

type Router interface {
	Route(method string, path string, handleFunc HandleFunc)
}
