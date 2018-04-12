package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type MyReader struct{}

func (read MyReader) Read(output []byte) (int, error) {
	fmt.Println("Help2")
	a := "A"
	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		output[i] = []byte(a)[0]
		fmt.Println("Help")
	}
	return 4, nil
}

func main() {
	M := MyReader{}
	stuff, _ := ioutil.ReadAll(M)
	log.Printf("%s", stuff)
}
