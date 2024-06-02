//go:build ignore
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/if_ether.h>

// global_map: map structure to receive user defined ports
// from user-space
struct {
    __uint(type, BPF_MAP_TYPE_ARRAY);
    __type(key, __u32);
    __type(value, __be16);
    __uint(max_entries, 1);
} global_data SEC(".maps");

SEC("xdp_prog")
// xdp_tcp_drop: Drops tcp packets targeting a specific
// port defined in global_data
int xdp_tcp_drop(struct xdp_md *ctx) {
    __be16 port = 4000;         // default port
    __be16 *port_ptr = &port;
    __u32 port_to_drop = 0;     // global_data index for port

    // find userdefined port in global_data
    port_ptr = bpf_map_lookup_elem(&global_data, &port_to_drop);
    if(port_ptr) { // checking if the port value exists
        bpf_printk("Value of port_ptr: %u\n", port_ptr);
        void *data = (void *)(long)ctx->data;           // packet begining
        void *data_end = (void *)(long)ctx->data_end;   // packet ending

        // retrieve eth, ip and tcp packets from packet data
        struct ethhdr *eth = data;
        if ((void*)eth + sizeof(*eth) <= data_end) {
            struct iphdr *ip = data + sizeof(*eth);
            if ((void*)ip + sizeof(*ip) <= data_end) {
                // checking for tcp connection
                if (ip->protocol == IPPROTO_TCP) {
                    struct tcphdr *tcp = (void*)ip + sizeof(*ip);
                    if (
                        (void*)tcp + sizeof(*tcp) <= data_end
                        &&
                        // compare tcp destination to user defined port
                        bpf_ntohs(*port_ptr) == tcp->dest
                    ) {
                        bpf_printk("%d", *port_ptr);
                        return XDP_DROP; // drop packet
                    }
                }
            }
        }
    }
    bpf_printk("Port not detected, passind xdp\n");
    return XDP_PASS; // pass packet
}

// attaching license to pass through verification
char __license[] SEC("license") = "Dual MIT/GPL";
