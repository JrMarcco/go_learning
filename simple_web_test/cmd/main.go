package main

import (
	"go_learning/simple_web/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}

	server.ListenAndServe()
}
