package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Make a channel that accepts type string
	c := make(chan string)

	// Loop over the links to check
	for _, link := range links {
		go checkLink(link, c)
	}

	// Loop the checks the channel for a message based on the number
	// of sites in the links slice.
	// NOTE: The call to receive a message from a channel is a blocking call
	for i := 0 ; i < len(links); i++ {
		fmt.Println(<- c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep, it's up"
}