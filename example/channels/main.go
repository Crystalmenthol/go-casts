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

	for i := 0; i < len(urls); i++ {
		fmt.Println(<- c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		log.Panic(err)
		c <- "Might be down I think"
		return 
	}
	c <- "Yep its up!"
	fmt.Println(link, "is up!")

}