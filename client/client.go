package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	url := "http://127.0.0.1:8081"
	headers := "A"
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 1024,
		},
		Timeout: 5 * time.Second,
	}

	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {
		request, err := http.NewRequest("GET", url, nil)
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
			fmt.Printf("Response: %+v\n\n", resp.Body())
			resp.Body.Close()
		}
	}
}
