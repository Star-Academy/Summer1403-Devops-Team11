package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
	arg := os.Args[1]
	fmt.Println(arg)
	protocol := "icmp"
	netaddr, _ := net.ResolveIPAddr("ip4", arg)
	fmt.Println(netaddr)
	conn, _ := net.ListenIP("ip4:"+protocol, netaddr)

	buf := make([]byte, 1024)
	numRead, _, _ := conn.ReadFrom(buf)
	fmt.Printf("% X\n", buf[:numRead])
}
