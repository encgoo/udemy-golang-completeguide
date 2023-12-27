# Mics

Several other topics about golang

## while loop
For is go's "while".
```go
    for _condition_ {

    }

    // forever loop
    for {

    }
```
Refer to [testFor.go](./testFor.go).

## channel definition
Even though the following declare a channel, but the channel is nil, and it can't be used later.
```go
    var c chan string
```
Always use make
```go
    c := make(chan string)
```

## non-buffered channel
For a single (non-buffered) channel, write and read must alternate. The read is waiting for the write before it can go on. More importantly, the write is also waiting for the read. This becomes tricky if the main routine use a channel to send information to a child routine. It has to spawn the child routine _first_ before writing. If it writes first, then it will hang there waiting for read before it can spawn a child rountine. This is a deadlock.

It seems more natural to let the child rountine to write to the channel and the main rountine to read. For information from main -> channel, it can just be normal function variable. 

## buffered channel
```go
    bufferedChan := make(chan string, 10)
```
For a buffered chan, the write can keep one going until the buffer is full. Refere to [channelpool.go](./channelpool.go).