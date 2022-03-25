package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func main(){
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	
	lw := logWriter{}
	// Example of how to read using the Writer interface with io.Copy
	// func Copy(dst io.Writer, src io.Reader) (written int64, err error)
	//io.Copy(os.Stdout, resp.Body) // <- we can substitute os.Stdout for our new function that implements the Writer interface
	io.Copy(lw, resp.Body)
	
}
// Example how to read implementing our own struct into the Write interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("We just wrote", len(bs), "bytes.")
	return len(bs), nil
}