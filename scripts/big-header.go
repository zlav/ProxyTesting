package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "http://my_first_url.vcap.me:8081"
	headers := "Authorization: Basic hi
there"
	// for i := 1; i <= 100000; i++ {
	// 	headers = headers + fmt.Sprint("A\n")
	// 	fmt.Printf("%d\n", i)
	// }
	// bigFile, err := ioutil.ReadFile("/Users/pivotal/Downloads/big-ass-picture.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// headers = string(bigFile)

	for i := 1; i <= 100000; i++ {
		request, _ := http.NewRequest("GET", url, nil)
		request.Header.Set("name", headers)
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			fmt.Printf("\nError was returned %+v\n\n", err)
			continue
		} else {
			fmt.Printf("Got Response: %+v\n", resp)
			// fmt.Printf("|")
		}
		resp.Body.Close()
	}

}
