package main

import (
	"fmt"
	"net/http"
)

type Reader struct{}

func NewReader() *Reader {
	return &Reader{}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	// fmt.Println([]byte(r.read))
	// for i, b := range []byte(r.read) {
	// 	p[i] = b
	// }
	// r.done = true
	// return len(r.read), nil
}

func main() {

	// url := "http://my_first_url.vcap.me:8081"
	url := "http://localhost:8081"
	headers := "A"
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1024,
		},
		//	Timeout: 5 * time.Second,
	}

	// ticker := time.NewTicker(1 * time.Second)
	//for range ticker.C {
	request, err := http.NewRequest("GET", url, NewReader())
	if err != nil {
		panic(err)
	}

	request.Header.Set("name", headers)
	resp, err := client.Do(request)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf("Error found: %+v\nResponse: %+v\n", err, resp)
		return
	} else {
		fmt.Printf("Response: %+v\n\n", resp)
		resp.Body.Close()
	}
	// }
}
