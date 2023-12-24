package main

import (
	"fmt"
	"math/rand" // random number
)

func main() {

	// need to create random numbers
	intArr := [11]int{rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100),
		rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100), rand.Intn(100)}

	for _, num := range intArr {
		if num%2 == 0 {
			fmt.Print("even number: ")
			fmt.Println(num)
		} else {
			fmt.Print("odd number: ")
			fmt.Println(num)
		}
	}
}
