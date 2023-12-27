package main

import (
	"fmt"
	"net/http"
	"time"
	_ "time"
)

var servers = [5]string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
	"http://amazon.com",
}

func main() {

	// Test function literal in the same thread

	names := [3]string{"apple", "orange", "pear"}

	for _, name := range names {
		// name is a varaible on range scope. Use it in a nested function
		var printName = func() {
			// a nested function can refer to the variable outside
			// No warning?
			fmt.Println(name)
		}

		// Here there is the same problem as below, but the compiler fails
		// to catch. Because, if we don't use go, it is ok.
		go printName()

		// function literal
		go func() {
			// This one is saying loop variable name captured by func literal
			// The problem here is that the name varaible is exactly the same
			// of the loop varialbe, and they are shared by two threads!
			fmt.Println(name)
			// If we remove "go" key word, then it is fine, because it is the
			// same thread using the same variable.
		}()
	}

	c := make(chan string)
	//var c chan int

	for _, link := range servers {
		go checkLink(link, c)
	}

	for l := range c {
		// The sleep here is on the main routine, which forces all children to wait
		// as well.
		// time.Sleep(time.Second)

		// So instead of this
		//go checkLink(l, c)

		// we use function literal (lambda function)
		// This one has error. The l below is a reference? not a value copy?
		// So when when the child rountine is using l, the main rountine might
		// go to the next iteration and change l.

		// go func() {
		// 	time.Sleep(5 * time.Second)
		// 	go checkLink(l, c)
		// }()

		// Instead, we shall make a copy when calling the function literal
		go func(link string) {
			time.Sleep(5 * time.Second)
			//fmt.Println(link)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// waiting here is better. It only affacts a child rountine.
	// time.Sleep(time.Second).
	// But it is not very logical, since checkLink is not supposed to wait
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Site: " + link + " is down!")
		c <- link
		return
	}
	fmt.Println("Site: " + link + " is up!")
	c <- link
}
