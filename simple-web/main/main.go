package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "welcome %s", r.URL.Path[1:]); err != nil {
			log.Fatalf("error %+v", err)
		}
	})

	log.Fatal(
		http.ListenAndServe(":8080", nil),
	)
}
