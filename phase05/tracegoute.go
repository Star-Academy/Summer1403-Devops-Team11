package main

import (
	// "encoding/binary"
	"fmt"
	"net"
	"os"
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

	icmpMsg := make([]byte, 8)

	icmpMsg[0] = 8
	icmpMsg[1] = 0
	icmpMsg[2] = 0
	icmpMsg[3] = 0
	icmpMsg[4] = 0
	icmpMsg[5] = 1
	icmpMsg[6] = 0
	icmpMsg[7] = 2

	_, err = conn.Write(icmpMsg)

	if err != nil {
		fmt.Println(err)
	}

	// reply := make([]byte, 1024)

	// n, err := conn.Read(reply)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(err)

}

// package main

// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"time"

// 	"golang.org/x/net/icmp"
// 	"golang.org/x/net/ipv4"
// )

// func main() {
// 	// Target IP address (change as needed)
// 	target := "8.8.8.8"

// 	// Resolve the IP address
// 	ipAddr, err := net.ResolveIPAddr("ip4", target)
// 	if err != nil {
// 		fmt.Println("Error resolving IP address:", err)
// 		return
// 	}

// 	// Create a raw socket for ICMP
// 	conn, err := net.DialIP("ip4:icmp", nil, ipAddr)
// 	if err != nil {
// 		fmt.Println("Error creating ICMP connection:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Create the ICMP Echo Request message
// 	msg := icmp.Echo{
// 		ID:   os.Getpid() & 0xffff, // Identifier
// 		Seq:  1,                    // Sequence number
// 		Data: []byte("Hello!"),     // Data to send
// 	}

// 	// Marshal the message into a byte slice
// 	b, err := msg.Marshal(nil)
// 	if err != nil {
// 		fmt.Println("Error marshaling message:", err)
// 		return
// 	}

// 	// Set a deadline for the response
// 	conn.SetDeadline(time.Now().Add(2 * time.Second))

// 	// Send the ICMP message
// 	if _, err := conn.Write(b); err != nil {
// 		fmt.Println("Error sending message:", err)
// 		return
// 	}

// 	// Prepare to receive a reply
// 	reply := make([]byte, 512)
// 	n, addr, err := conn.ReadFrom(reply)
// 	if err != nil {
// 		fmt.Println("Error reading reply:", err)
// 		return
// 	}

// 	// Parse the ICMP reply message
// 	r, err := icmp.ParseMessage(icmp.ProtocolICMP, reply[:n])
// 	if err != nil {
// 		fmt.Println("Error parsing reply:", err)
// 		return
// 	}

// 	// Check the type of the reply
// 	switch r.Type {
// 	case ipv4.ICMPTimeExceeded:
// 		fmt.Printf("Received Time Exceeded from %s\n", addr.String())
// 	case ipv4.ICMPReply:
// 		fmt.Printf("Received Reply from %s\n", addr.String())
// 	default:
// 		fmt.Printf("Unexpected reply from %s: %v\n", addr.String(), r)
// 	}
// }
