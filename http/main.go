package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type LogWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	lw := LogWriter{}
	io.Copy(lw, resp.Body)
}

func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(fmt.Sprintf("Wrote %d bytes \n", len(bs)))

	return len(bs), nil
}
