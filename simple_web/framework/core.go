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
	router map[string]*TrieTree
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

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router[GET].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router[POST].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router[PUT].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router[DELETE].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx := NewContext(request, writer)
	router := c.FindRouter(request)
	if router == nil {
		ctx.NotFoundJson("Not Found")
		return
	}

	if err := router(ctx); err != nil {
		ctx.ErrJson("Inner Error")
		return
	}
}

func (c *Core) FindRouter(request *http.Request) ControllerHandler {
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
