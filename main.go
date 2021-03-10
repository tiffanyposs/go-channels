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
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // make a new channel with values of type string

	for _, link := range links {
		go checkLink(link, c) // pass channel
	}

	// wait for the channel to return some value
	// when it does, assign it to the variable l (link)
	for l := range c {
		// Function literal - (anonymous function)
		go func(link string) {
			time.Sleep(5*time.Second) // delay 5 seconds
			checkLink(link, c) // wait for message from c, and call checkLink with the value
		}(l) // must pass l to copy it, prevents funcy behavior
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking call
	if err != nil {
		msg := "might be down!"
		fmt.Println(link, msg)
		c <- link // send the link to the channel
		return
	}

	msg := "is up!"
	fmt.Println(link, msg)
	c <- link // send the link to the channel
}
