package main

import (
	"fmt"
	"reflect" //check type
)

var pi float32 = 3.15

func main() {
	var cards string = "Aces"

	if pi == 3.14 {
		fmt.Println(cards)
	}

	arrayValue()
	sliceOrArray()
	sliceAppend()
}

func arrayValue() {
	fmt.Println("* arrayValue:")
	var arrInt = [5]int{1, 2, 3, 4, 5}
	// passing a copy
	modifyArray(arrInt)
	// no impact
	fmt.Println(arrInt[0])

	// passing a pointer just like c++
	modify(&arrInt)
	// Now changed
	fmt.Println(arrInt[0])

	// for loop
}

// input argument must state size, because it is taking an array. NOT a slice
// (arrSlce []int) on the other hand will take a slice.
func modifyArray(arr [5]int) {
	arr[0] = 100
}

// input argument still need to state size, different from C/C++. [5]int
// together is a type
func modify(arr *[5]int) {
	// This will show a ptr as type!
	fmt.Println(reflect.TypeOf(arr).Kind())
	arr[0] = 100
}

func sliceAppend() {
	fmt.Println("* sliceAppend: ")
	var arr = [6]int{1, 2, 3, 4, 5, 6}

	var slce []int = arr[2:5]

	fmt.Println(reflect.TypeOf(slce).Kind())

	fmt.Print("before modifying array: ")
	fmt.Println(slce[0])
	arr[2] = 6
	// This shows slce is a reference of arr _at_this_point_
	// Modifying arr will affect slce
	fmt.Print("after modifying array: ")
	fmt.Println(slce[0])
	// Now append.
	fmt.Print("Before append, arr[5]: ")
	fmt.Println(arr[5])
	// after append a new slice is created.
	newslce := append(slce, 10)
	fmt.Print("After append, arr[5]: ")
	fmt.Println(arr[5])
	// slce remains size 3
	fmt.Println(len(slce))
	// newslce has size 4
	fmt.Println(len(newslce))
	// newslce and slce are still referencing to arr at this point
	arr[2] = 100
	fmt.Println(slce[0])
	fmt.Println(newslce[0])

	// Now slce exceeds the capacity of arr. slce now references to a new arr
	slce = append(slce, 10, 20, 30, 1000)
	// changing the original arr doesn't affect slce any more.
	arr[2] = 0
	fmt.Println(slce[0])

}

func sliceOrArray() {
	fmt.Println("* sliceOrArray: ")
	// Array
	var arr = [3]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(arr).Kind())

	// Slice. Note that the referenced array is no longer accessible??
	var slce = []int{1, 2, 3}
	fmt.Println(reflect.TypeOf(slce).Kind())

	// Array
	var tmp = [...]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(tmp).Kind())
}
