package main

import (
	"fmt"
	"time"
)

const NumThreads = 3

func channelPool() {

	chPool := make(chan int, NumThreads)

	for i := 1; i < 10; i++ {
		// if the pool is full, this will hold
		chPool <- i
		go prInt(i, chPool)
	}
}

// function takes chan not buffered chan
func prInt(i int, c chan int) {
	fmt.Println(i)
	time.Sleep(2 * time.Second)
	fmt.Print("Done")
	fmt.Println(<-c)
}
