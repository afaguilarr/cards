package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	webPages := []string{
		"http://google.com",
		"http://golang.org",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
	}
	arePagesAlive(webPages)
}

func arePagesAlive(webPages []string) {
	c := make(chan string)
	for _, webPage := range webPages {
		go isWebPageAlive(webPage, c)
	}
	for webPage := range c {
		go func(webPage string) {
			time.Sleep(2 * time.Second)
			isWebPageAlive(webPage, c)
		}(webPage)
	}
}

func isWebPageAlive(webPage string, c chan string) {
	resp, err := http.Get(webPage)
	if err != nil {
		fmt.Printf("Error while getting the %s GET endpoint: %s\n", webPage, err)
		c <- webPage
		log.Fatalf("Error while getting the %s GET endpoint: %s\n", webPage, err)
	}
	isAlive(webPage, resp.StatusCode)
	c <- webPage
}

func isAlive(webPage string, status int) {
	if status == 200 {
		fmt.Printf("The %s GET endpoint is healthy, status code: %v\n", webPage, status)
		return
	}
	fmt.Printf("The %s GET endpoint is not healthy, status code: %v\n", webPage, status)
}
