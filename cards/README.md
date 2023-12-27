# Section 3: Deeper into Go

## variable initialization
Inside a function these two ways are the same
```go
a := 3.14
var b float64 = 3.14
```
But outside a function, only the second one is valid.

## Array data type
Examples:
```go
var arr[10] int
// This is an array
intArr := [5]int{1,2,3,4,5}

// This is an array
var intArr = [...]{1,2,3,4,5,6,7}
```
We can also init part of the array. Here only the third one (0 based) is inited, others remain empty strings.  
```go
var keyValArr = [5]string{2:"Second one"}
```

### function argumnent as copy
Array can be passed as a copy or a pointer, similar to C/C++. 

### loop through an array
```go
for i, item := range items {
    // Need to use both variables.
}
```

## Slice
Slice is a reference to a contiguous segment of an aray. So slice is just a view on an array. A slice just needs to keep the start pointer, the length, and the full capacity of a referenced array. 
```go
// This is a slice
intArr := []int{1,2,3,4,5,6}

// This is an array
var intArr = [3]{1,2,3}
// This is a slice
slce := intArr[1:2]
// This is also a slice
var newslce[] int = intArr[:]

```

### in place operation
Split() returns a slice still points to the original data, _not_ a copy.

### append
A slice is a reference of (part of) an array. When append, if it is still within the capacity of the original array, the original array will be modified. See the example of sliceAppend in [main.go](./main.go). But if the append exceeds the capacity, a new array will be created, and then the new slice returned by the append function call is no longer coupled to the original array. 

The ambiguity whether the slice is still coupled to the original array can be fixed by always calling a copy function to make a copy. 

### Joining strings
There is a package called strings. It has a Join function that can join string slices.
```go
import "strings"

strings.Join
```
### slice as argument
When others are used as function arguments, golang uses pass by value. So it is a copy that is passed into the function.

But slice is different. When a slice is passed as an argument, it is passed by reference. The same copy is used. 
]
## Type with methods
This is almost like a classs in C++.
```go
// member variables go into the the struct of the class_name
type class_name struct {
    privateMemVar   string
    PublicMemVar    int
}

// member functions here. They are also referred as receiver functions
func (obj class_name) privateFunc() {

}

func (obj class_name) PublicFunc() {

}
```
Check embedded struct for inheritance/composition.

## Returning multiple values
A function can return multiple values.
```go
func main(){
    int1, int2 := goo()
}

func goo() (int, int) {
    return 1, 2
}
```

## Byte slices
Byte is char in C/C++, which represents a ASCII char. rune is for UTF-8. Byte is stored as ASCII value. If we cast a string into byte slice and print, it prints out the ASCII values. 

## Swap assign
```go
// this will swap the values of a and b, just like python
a, b = b, a
```

## Testing
* Use filename_test.go for a test file for filename.go
* Use TestFuncName for a function to test FuncName
* Use t.Errorf to output error message when an error is detected.
* Use `go test -timeout Xs` to set a timeout for a go test.

## Reference type and value type
Reference types include:
* slice
* map
* channel
* interface
* function type

## while loop for ever
Use for
```go 
for {

}

```