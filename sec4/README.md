# Organizing data with structs

NOTE the parent folder of this module contains multiple modules. Make sure `go work with sec4` has been executed. Check go.work above to verify. Otherwise the import doesn't work.

## Struct 
### Definition
```go
type person struct {
    firstName string
    lastName string
}
```
### Declaration
```go
    alex := person{"Alex", "Anderson"}
    // or dict like
    alex := person{firstName: "Alex", lastName: "Anderson"}
    // empty init string->"", numbers ->0, bool ->false
    var alex person
```

### Embedding Structs
Embedding structs is how golang handles composition. A struct composes another struct with definitions. 

### Receiver function
Receiver function on value or receiver function on a pointer. Only the latter one can modify the values/properties. The first one is passed by value. So the change won't be "returned". 
#### Pointer shortcut
Go supports pointer shortcut for simplicity. The receiver function is defined on a pointer of a struct, but when it is called, instead of using a pointer, we can use the value directly. Go will convert it to pointer as a shortcut.
