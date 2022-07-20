package main

import (
	"go_learning/simple_web/framework"
	"go_learning/simple_web_test/internal"
	"net/http"
)

func main() {
	core := framework.NewCore()
	internal.RegisterRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    ":8080",
	}

	server.ListenAndServe()
}
