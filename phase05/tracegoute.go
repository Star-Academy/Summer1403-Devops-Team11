package main

import (
	// "encoding/binary"
	"fmt"
	"net"
	"os"

	// "syscall"
	// "encoding/binary"
	"golang.org/x/net/ipv4"
	//"time"
)

type HeaderFlags int

const (
	MoreFragments HeaderFlags = 1 << iota // more fragments flag
	DontFragment                          // don't fragment flag
)

type Header struct {
	Version  int         // protocol version
	Len      int         // header length
	TOS      int         // type-of-service
	TotalLen int         // packet total length
	ID       int         // identification
	Flags    HeaderFlags // flags
	FragOff  int         // fragment offset
	TTL      int         // time-to-live
	Protocol int         // next protocol
	Checksum int         // checksum
	Src      net.IP      // source address
	Dst      net.IP      // destination address
	Options  []byte      // options, extension headers
}

func checksum(msg []byte) uint16 {
	sum := 0
	for i := 0; i < len(msg)-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if len(msg)%2 == 1 {
		sum += int(msg[len(msg)-1]) * 256
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	return uint16(^sum)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host\n", os.Args[0])
		os.Exit(1)
	}

	host := os.Args[1]

	// Resolve IP address
	ipAddr, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		fmt.Println("Error  resolving IP address:", err)
		os.Exit(1)
	}

	conn, err := net.DialIP("ip4:icmp", nil, ipAddr)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		os.Exit(1)
	}

	ipv4Conn := ipv4.NewConn(conn)

	err = ipv4Conn.SetTTL(64)

	icmpMsg := make([]byte, 8)

	icmpMsg[0] = 8
	icmpMsg[1] = 0
	icmpMsg[2] = 0
	icmpMsg[3] = 0
	icmpMsg[4] = 0
	icmpMsg[5] = 1
	icmpMsg[6] = 0
	icmpMsg[7] = 2

	// icmpMsg := Header{
	// 	Version:  4,                        // Example for IPv4
	// 	Len:      20,                       // Example header length in bytes
	// 	TOS:      0,                        // Default type-of-service
	// 	TotalLen: 40,                       // Example total length
	// 	ID:       54321,                    // Example identification
	// 	Flags:    0,                        // Example flag value
	// 	FragOff:  0,                        // No fragmentation
	// 	TTL:      64,                       // Example TTL value
	// 	Protocol: 6,                        // Example for TCP
	// 	Checksum: 0,                        // Placeholder for checksum
	// 	Src:      net.ParseIP("127.0.0.1"), // Example source address
	// 	Dst:      net.ParseIP("127.0.0.1"), // Example destination address
	// 	Options:  []byte{0x01, 0x02},       // Example options
	// }
	// data, err := icmpMsg.ToBytes()

	checksum := checksum(icmpMsg)
	icmpMsg[2] = byte(checksum >> 8)
	icmpMsg[3] = byte(checksum)

	_, err = conn.Write(icmpMsg)

	if err != nil {
		fmt.Println(err)
	}

	buff := make([]byte, 512)

	_, err = conn.Read(buff)

	fmt.Println(buff)
}
