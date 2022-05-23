package main

import (
	"fmt"
	"go_learning/simple-web/pkg/server"
	"log"
	"net/http"
)

func main() {
	simpleHttpServer := server.NewSimpleHttpServer(":8080")

	simpleHttpServer.Route("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "welcome %s", r.URL.Path[1:]); err != nil {
			log.Fatalf("error %+v", err)
		}
	})

	if err := simpleHttpServer.Start(); err != nil {
		log.Fatalf("%-v", err)
	}
}
