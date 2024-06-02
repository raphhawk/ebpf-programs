# eBPF_drop: A program to drop TCP packets from a network device on a specified port

The tool uses libbpf and golang on userspace.</br>

The tool can take user configurable port and network device.</br>

## Usage
```bash
    # Install dependencies for void linux
    sudo xbps-install libbpf-devel clang linux-headers llvm make elfutils-devel bpftool bpftrace netcat go 

    # from the root of project execute the following 
    # To build the ebpf kernel program and userspace program
    make build
    
    # To run the program with default port i,e (4000) and default network i,e (loopback)
    make run
    
    # To run with user specified port/nw device/both 
    sudo ./builds/epbf_drop -port=8080 -nw=lo

    # To clean all generated elf/go binaries
    make clean
```
