package main

import (
	"io"
	"io/ioutil"
	"log"
)

type Reader struct {
	read string
	done bool
}

func newReader(toRead string) *Reader {
	return &Reader{toRead, false}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.done {
		return 0, io.EOF
	}
	for i, b := range []byte(r.read) {
		p[i] = b
	}
	r.done = true
	return len(r.read), nil
}

func main() {
	reader := newReader("test")
	out, _ := ioutil.ReadAll(reader)
	log.Printf("%s", out)
}
