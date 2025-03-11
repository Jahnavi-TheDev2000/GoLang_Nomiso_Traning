package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom writer
type logWriter struct {
}

func main() {
	//func Get(url string) (resp *Response, err error)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	// fmt.Println(resp)
	// fmt.Printf("Response is of Type: %T\n", resp)

	//1 way
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs)) //it is giving the html response body

	//replace the upper code
	//2 way
	io.Copy(os.Stdout, resp.Body)
}

// custom implemntation of writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Byte of data:", len(bs))
	return len(bs), nil
}
