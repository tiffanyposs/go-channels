package main

import (
	"fmt"
	"net/http"
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
		go checkLink(link, c)
	}

	fmt.Println("main", <-c)  // waiting for a message from c
	fmt.Println("main2", <-c) // waiting for a message from c
	fmt.Println("main3", <-c) // waiting for a message from c
	fmt.Println("main4", <-c) // waiting for a message from c
	fmt.Println("main5", <-c) // waiting for a message from c
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) // blocking call
	if err != nil {
		msg := "might be down!"
		fmt.Println(link, msg)
		c <- msg // send the message to the channel
		return
	}

	msg := "is up!"
	fmt.Println(link, "is up!")
	c <- msg // send the message to the channel
}
