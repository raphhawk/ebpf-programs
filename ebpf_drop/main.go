package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// log levels
const (
	e = "[ERROR]"
	i = "[INFO]"
	d = "[DEBUG]"
)

func main() {
	// Set flags
	port := flag.Int("port", 4000, "get user specific port")
	nw := flag.String("nw", "lo", "get netword device")
	flag.Parse()

	// For kernel version < 5.11
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	// Set kernel objects
	var obj tcp_dropObjects
	if err := loadTcp_dropObjects(&obj, nil); err != nil {
		log.Fatal(e, " Loading epbf objects:", err)
	}
	defer obj.Close()

	// Update global_map with the respective ports in kernel program
	if err := obj.GlobalData.Update(uint32(0), uint16(*port), 0); err != nil {
		log.Fatal(e, " Updating ePBF map:", err)
	}
	log.Println(i, "ePBF Port data updated successfully to :", *port)

	// Set Network Interface
	ifname := *nw
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatalf("%v Getting interface %s: %s", e, ifname, err)
	}
	log.Printf("%v Setting network device as %v\n", i, ifname)

	// Attach kernel program to XDP event
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   obj.XdpTcpDrop,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatal(e, " Attaching XDP:", err)
	}
	defer link.Close()
	log.Printf("%v All TCP packets to the port %v from network %v will be dropped\n", i, *port, *nw)

	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt) // listen for SIGINT
	select {
	case <-stop:
		log.Println(i, " Received Interrupt signal, exiting...")
		return
	}
}
