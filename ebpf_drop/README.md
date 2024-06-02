# eBPF_drop: A program to drop TCP packets from a network device on a specified port

This tool uses libbpf and golang on userspace.

```make
    // To build the ebpf kernel program and userspace program
    make build
    
    // To run the program with default port i,e (4000) and default network i,e (loopback)
    make run

    // To clean all generated elf/go binaries
    make clean
```
