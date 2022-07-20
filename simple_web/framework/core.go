package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Println("### ServeHTTP ###")
	ctx := NewContext(request, writer)

	if router, ok := c.router["foo"]; ok {
		if err := router(ctx); err != nil {
			return
		}
	}
	return
}
