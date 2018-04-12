package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// fmt.Printf("%+v\n", time.Now())
	// time.Sleep(1 * time.Second)

	// // for i := 0; i < 20; i++ {
	// nBytesRead, err := r.Body.Read(readOut)
	// if err != nil {
	// 	fmt.Printf("Error reading %s", err)
	// 	return
	// }
	// fmt.Printf("read %d bytes\n", nBytesRead)
	// fmt.Printf("%s\n", string(readOut[0:nBytesRead]))

	// }
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

	if err := conn.Close(); err != nil {
		fmt.Printf("error closing: %s", err)
		return
	}
	fmt.Println("I CLOSED IT!")

	// err = r.Body.Close()
	// if err != nil {
	// 	fmt.Printf("error closing: %s\n", err)
	// 	return
	// }
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
