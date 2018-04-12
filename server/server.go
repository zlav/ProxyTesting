package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// fmt.Printf("%+v\n", time.Now())
	// time.Sleep(1 * time.Second)
	fmt.Printf("|")
}

func main() {

	// http.HandleFunc("/", handler)
	s := &http.Server{
		Addr:    "localhost:8081",
		Handler: http.Handler(&handler{}),
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 2,
	}

	log.Fatal(s.ListenAndServe())
}
