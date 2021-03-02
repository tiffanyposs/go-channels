# Channels in Go

Below example waits for each link's `checkLink` to finish before it continues.


```go

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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // blocking call
	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}
	fmt.Println(link, "is up!")
}

```

## Go Routines

* Go Routine - Engine that executes code

The keyword `go` is used in front of blocking calls to create a Go Routine. Doing so allows a blocking call to pass the execution back to the main program.