package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
	}
	fmt.Println(colors)

	var colrs map[string]string
	fmt.Println(colrs)

	clors := make(map[string]string)
	clors["red"] = "#ff0000"
	fmt.Println(clors)

	// delete a key-value pair from a map
	delete(clors, "red")

	// Iterate over a map
	for clr, hex := range colors {
		fmt.Printf("%s is %s\n", clr, hex)
	}
}
