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

By default, Go tries to use one CPU core on your machine. In this case, it simply allows you to run other code while your program waits for an asynchronous call. Once the async call resolves the Go routine will pass the work back to the methe main thread.

* `concurrencey` - we can have multiple threads executing code. If one thread is blocks, another one is picked up and worked on.
* `parallelism` - multiple threads executing at the same time. Requires multiple CPUs.

A Go Routine has a main routine and child routines.

## Channels

Channels are used to communicate between Go routines.

Declare a channel:

```go
  // create new channel with chan keyword
  c := make(chan string)

  ...

  // pass it into the function
  func myFunc(link string, c chan string) {
    // do stuff
  }
```

Syntax:

* `channel <- 5` - Send value 5 into this channel
* `myNumber <- channel` - Wait for value to be sent to the channel. When we get one, assign the value to `myNumber`
* `fmt.Println(<- channel)` - Wait for the value to be sent into the channel. When we get one, log it out immediately.