package main

import (
	"flag"
	"log"
	"net"

	"github.com/cilium/ebpf/link"
)

const (
	e = "[ERROR]"
	i = "[INFO]"
	d = "[DEBUG]"
)

func main() {
	port := flag.Int("port", 4000, "get user specific port")
	nw := flag.String("nw", "lo", "get netword device")
	flag.Parse()

	var obj tcp_dropObjects
	if err := loadTcp_dropObjects(&obj, nil); err != nil {
		log.Fatal(e, " Loading epbf objects:", err)
	}
	defer obj.Close()

	if err := obj.GlobalData.Update(uint32(0), uint16(*port), 0); err != nil {
		log.Fatal(e, " Updating ePBF map:", err)
	}
	log.Println(i, "ePBF Port data updated successfully to :", *port)

	ifname := *nw
	iface, err := net.InterfaceByName(ifname)

	if err != nil {
		log.Fatalf("%v Getting interface %s: %s", e, ifname, err)
	}

	link, err := link.AttachXDP(link.XDPOptions{
		Program:   obj.XdpTcpDrop,
		Interface: iface.Index,
	})

	if err != nil {
		log.Fatal(e, " Attaching XDP:", err)
	}

	log.Printf("%v Setting network device as %v\n", i, ifname)
	log.Printf("%v All TCP packets to the port %v from network %v will be dropped\n", i, *port, *nw)
	defer link.Close()
	select {}
	/*
		tick := time.Tick(time.Second)
		stop := make(chan os.Signal, 5)
		signal.Notify(stop, os.Interrupt)

		for {
			select {
			case <-tick:
				//log.Printf("Received %d packets")
			case <-stop:
				log.Println("Received signal, exiting...")
				return
			}
		}
	*/
}
