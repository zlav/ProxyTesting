package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request %+v", r)
	fmt.Printf("I'm going to sleep for a second\n")
	time.Sleep(1 * time.Second)

	fmt.Printf("writing response status code\n")
	w.WriteHeader(420)
	fmt.Printf("writing response body\n")
	w.Write([]byte("thats too many bytes, please go away."))

	fmt.Printf("Now I'm going to close the connection\n")
	hj, ok := w.(http.Hijacker)
	if !ok {
		fmt.Println("not a hijacker!")
		return
	}

	// Hijack the rw.
	conn, _, err := hj.Hijack()
	if err != nil {
		fmt.Printf("error hijacking: %s", err)
		return
	}

	// err = r.Body.Close()
	// if err != nil {
	// 	fmt.Printf("error closing body: %s\n", err)
	// 	return
	// }

	if err := conn.Close(); err != nil {
		fmt.Printf("error closing: %s", err)
		return
	}
	fmt.Println("I CLOSED IT!")
}

func main() {

	s := &http.Server{
		// Addr:		"localhost:8081"
		// Gorouter setting
		Addr:    "localhost:4567",
		Handler: http.Handler(&handler{}),
	}
	fmt.Printf("Serving on %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
