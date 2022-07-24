package framework

type HandlerFunc func(ctx *Context) error

type HandlerChain []HandlerFunc
