package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
		"http://taskflowhr.com",
	}

	c := make(chan string)

	for _, link := range links {
		go ping(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			ping(link, c)
		}(l)
	}
}

func ping(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(fmt.Sprintf("%s might be down!", link))
		c <- link
		return
	}

	fmt.Println(fmt.Sprintf("%s is up!", link))
	c <- link
}
