package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkLink(url, c)
	}

	for {
		go checkLink(<- c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		log.Panic(err)
		c <- link
		return 
	}
	c <- link
	fmt.Println(link, "is up!")

}