package main

import (
	"fmt"
	"net/http"
)

var servers = [5]string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func main() {
	for _, link := range servers {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Site: " + link + " is down")
		return
	}
	fmt.Println("Site: " + link + " is up!")
}
