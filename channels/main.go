package main

import (
	"fmt"
	"net/http"
	"time"
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

	// Infinite loop that uses chanel messages to contniually
	// check if a website is up or down.
	//
	// Alternative syntax using the range keyword to check the
	// channel for messages.
	//
	// Use a function liteeral (temporary function) to wrap both the
	// sleep call and the check link call in a go routine
	// NOTE: The link to check must be passed into the function literal
	//       as an argumnet so the value of the link is not "captured"
	//       by the function literal from outside its scope
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

// Check website and send message to channel based on response
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// Sending the website checked as a channel message
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// Sending the website checked as a channel message
	c <- link
}