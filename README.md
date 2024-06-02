# Accuknox assesment solutions

## Problem statement 1: Drop packets using eBPF

### Write an eBPF code to drop the TCP packets on a port (def: 4040). Additionally, if you can make the port number configurable from the userspace, that will be a big plus.
The code and explanation is implemented inside `ebpf_drop` directory. Feel free to check it out.

## Explain what the following code is attempting to do?
```go
package main

import "fmt"

func main() {
    cnp := make(chan func(), 10)
    for i := 0; i < 4; i++ {
        go func() {
            for f := range cnp {
                f()
            }
        }()
    }
    cnp <- func() {
        fmt.Println("HERE1")
    }
    fmt.Println("Hello")
}
```
### Explaining how the highlighted constructs work?
- [x] The highlighted code consists of a buffered channel `cnp` that can take upto `10 functions` without blocking any resource.
- [x] It also consists of a `loop with 4 iterations`, with each iteration it spawns a new `go routine` that ranges over the `cnp` channel and executes the function passed by it.
- [x] it then passes a single function to the `cnp` channel that prints `"HERE1"` to stdout. 
- [x] it then prints `"Hello"` to stdout. 

### Giving use-cases of what these constructs could be used for?
- [x] As these constructs consists of 4 go routines simultaniously listening on a buffered channel that send functions, this code can be used for functions that can be `executed concurrently and parallely on highly available resources` effectively. 
- [x] This code can be used in various different usecases such as `implementing Event Handling, Pipeline Processing(ETL), Resource Pooling for Scalablity` etc.

### What is the significance of the for loop with 4 iterations?
- [x] Each iteration of the loop `spawns a new goroutine`, so there will be four goroutines running concurrently once the loop completes. These goroutines will continuously listen for function values sent to the cnp channel and execute them as they arrive.
- [x] Therefore, the significance of the for loop with 4 iterations is that it `enables parallelism and concurrent execution` of functions, potentially improving performance and responsiveness of the program, especially in scenarios where tasks can be executed independently.

### What is the significance of make(chan func(), 10)?
- [x] make(chan func(), 10) creates a `buffered channel` that allows for asynchronous communication between goroutines, prevents `goroutine blocking`, and provides control over concurrency and resource utilization.
- [x] the channel can take upto 10 functions without blocking any running resources. This `limits the rate of functions` and lets the go routines to work efficiently and safely. 

### Why is “HERE1” not getting printed?
- [x] The reason "HERE1" is not getting printed is because the sending of the function to the channel cnp happens before any of the goroutines have started to consume values from the channel.
- [x] We are not `waiting` for the spawned go routines to finish execution, our main routine exits more quickly than the spawned routines. Hence we cannot determine if the passed functions to the channel have been executed or not. To wait for the go routines to finish execution we must use `sync.Waitgroup or a done channel`.
