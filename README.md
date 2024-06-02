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

### Giving use-cases of what these constructs could be used for.

### What is the significance of the for loop with 4 iterations?

### What is the significance of make(chan func(), 10)?

### Why is “HERE1” not getting printed?
