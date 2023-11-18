package main

import (
	"io"
	"log"
	"os"
)

func readFile() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, f)
}
func main() {
	readFile()
}
