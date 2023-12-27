package main

import (
	"fmt"
	"time"
)

func main() {
	channelPool()
	testBufferChannel()
	testFor()
}

func testBufferChannel() {
	// first a non-buffered channel
	c := make(chan int)
	// If we don't have this one, c<-2 after c<-1 will hang
	go printInt(c)
	c <- 1
	// the following will hang the app, since none reads from the chan yet
	// it causes an error "all goroutines are asleep"
	//c <- 2

	fmt.Println("main thread unblocked to write again")
	go printInt(c)
	c <- 2
	// Note that 2 might not be printed here, since the main thread is done.

	cb := make(chan int, 3)
	// cb is a buffered channel, we can write twice before spawning
	cb <- 1
	cb <- 2
	go printInt(cb)
	cb <- 3
	go printInt(cb)
	go printInt(cb)
	time.Sleep(3 * time.Second)
}

func printInt(c chan int) {
	//time.Sleep(time.Second)
	// Now we read
	fmt.Println(<-c)
}
