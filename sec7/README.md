# Section 7: Channels and go routines

## Go routine
Put a `go` keyword in front of a function call to spawn out a new (child) thread. There is a go scheduler behind the scene and manages those threads. If there is only one CPU core, then it is time slicing. This is the _default_ case. If there are multiple cores, we can configure go to use more than one CPU. Concurrency is time-slicing. When one block is blicked, the other thread can run. Parallelism is real multi-core. 

## Channel
A channel is used to communicate between go routines. A channel is associated with a type. 

### declaration
Declare a chan like this:
```go
    cStr := make(chan string)
    var cInt chan int
```
### Use
The channel must be one of the argument of the function we want to run in a child routine. 
```go
func checkLink(link string, c chan string){

}
```

### Sending message
```go
cInt <- 3
```
### Reading message
```go
intIn <- cInt

for l := range cInt {

}
```

### Wait for a child routine
When a channel is reading, it will wait for the other side the write first. This is how we can control the main routine to wait for the child routine. 
There are several ways to do this.
* Create a single channel from main and share it with all the children. A child write into the channel when it is done. The main routine read _all_ message from the channel. By doing so, it will wait for all children to finish first.
* One channel per child?


### Repeating routine
When a child routine is done, it returns its signature/id, then the main rountine knows what to restart.

### Sleeping in a routine
```go
import "time"

time.Sleep(5*time.Second)
```

### Function literal
It is like python lamda, or Javascript anonymous function.

Call function literal in go rountine.
```go
go func () {
    time.Sleep(5*time.Second)
    // NOTE DON'T put go below to spawn a new thread. It is already in a child 
    //routine here
    checkLink(l, c)
} ()
```

### Nested function and function literal
Function literal can be considered as anonymous nested function. So they are simliar, in the sense that they both live inside another (parent) function. 
```go
    func foo(){
        str := "cool"
        // declaration and definition
        var displayString := func(){
            // use local variable of parent function directly
            fmt.Println(str)
        }
        // execute
        displayString()

    }
```
They can use the _local_ variables of the parent function _directly_. This means it is _not_ a copy. 

This in general is OK if there is only one thread running. But if the function literal or the nested function is running in another go rountine, then this is a data-coupling issue. The main thread can change the value when the child rountine is using it.

In summary, when a nested function or a functin literal is dispatched into a child routine, make sure no local variable is shared, except channels. Channels are ok, since read and write are locked. If the main routine already wrote to a channel, it can't even change it by writing it again until the child routine can have a chance to read it.  