package main

import "fmt"

func testFor() {
	sum := 0

	// For is like while
	for sum < 5 {
		sum += 1
	}
	fmt.Println(sum)

	// Just for is a for-ever loop. How does channel stop it?
	c := make(chan string)
	// This is a mere declaration, and the channel is never created.
	//var c chan int
	go otherThread("First", c)
	go otherThread("Second", c)

	for i := 0; i < 2; i++ {
		fmt.Println(<-c)
	}

	// This is also a for ever loop. It won't stop waiting, if none writes, then
	// it is a deadlock.
	// for l := range c {
	// 	fmt.Println(l)
	// }

	// Same deablock as above!
	// for {
	// 	// This is forever loop. So it will be a deadlock. Not a good thing to
	// 	// try.
	// 	fmt.Println(<-c)
	// }
	// fmt.Println("Exited for-loop")
}

func otherThread(str string, c chan string) {
	fmt.Println(str)
	c <- "Done: " + str
}
