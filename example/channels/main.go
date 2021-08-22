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

	for _, url := range urls {
		checkLink(url)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println(link, "is up!")

}