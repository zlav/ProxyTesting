package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

var count int

func Counter() int {
	count = count + 1
	return count
}

func (r *Reader) Read(p []byte) (n int, err error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("|")
	p[0] = 'A'
	if Counter() > 80 {
		return 0, io.EOF
	}
	return 1, nil
}

func main() {

	// url := "http://localhost:8081"
	// Gorouter Setting
	url := "http://my_first_url.localhost.routing.cf-app.com:8081"

	headers := "A"
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1024,
		},
	}

	fmt.Println(url)
	request, err := http.NewRequest("GET", url, NewReader())
	if err != nil {
		panic(err)
	}

	request.Header.Set("name", headers)
	resp, err := client.Do(request)

	if err != nil {
		fmt.Printf("Error found: %+v\nResponse: %+v\n", err, resp)
		return
	} else {
		defer resp.Body.Close()
		fmt.Printf("Response: %+v\n\n", resp)
	}
}
