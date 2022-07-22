package framework

import (
	"log"
	"net/http"
	"strings"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Core struct {
	router      map[string]*TrieTree
	middlewares []ControllerHandler
}

func NewCore() *Core {
	router := map[string]*TrieTree{}
	router[GET] = NewTrieTree()
	router[POST] = NewTrieTree()
	router[PUT] = NewTrieTree()
	router[DELETE] = NewTrieTree()

	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router[GET].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router[POST].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router[PUT].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handlers ...ControllerHandler) {
	all := append(c.middlewares, handlers...)
	if err := c.router[DELETE].AddRouter(url, all); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx := NewContext(request, writer)
	handlers := c.FindRouter(request)
	if handlers == nil {
		ctx.NotFoundJson("Not Found")
		return
	}

	ctx.SetHandlers(handlers)
	if err := ctx.Next(); err != nil {
		ctx.ErrJson("Inner Error")
		return
	}
}

func (c *Core) FindRouter(request *http.Request) []ControllerHandler {
	reqUrl := request.URL.Path
	reqMethod := strings.ToUpper(request.Method)

	if methodRouters, ok := c.router[reqMethod]; ok {
		return methodRouters.FindHandler(reqUrl)
	}
	return nil
}

func (c *Core) Group(prefix string) *Group {
	return NewGroup(c, prefix)
}
