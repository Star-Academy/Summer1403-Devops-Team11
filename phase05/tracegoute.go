package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const ProtocolICMP = 1

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host\n", os.Args[0])
		os.Exit(1)
	}

	host := os.Args[1]

	// Resolve IP address
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		fmt.Println("Error resolving IP address:", err)
		os.Exit(1)
	}

	for ttl := 1; ttl <= 30; ttl++ {
		// Create a raw socket
		conn, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
		if err != nil {
			fmt.Println("Error creating socket:", err)
			os.Exit(1)
		}

		defer conn.Close()

		if err := conn.IPv4PacketConn().SetTTL(ttl); err != nil {
			fmt.Println("Error setting TTL:", err)
			os.Exit(1)
		}

		// Create ICMP packet
		icmpMsg := make([]byte, 8)
		icmpMsg[0] = 8
		icmpMsg[1] = 0
		icmpMsg[2] = 0
		icmpMsg[3] = 0
		icmpMsg[4] = 0
		icmpMsg[5] = 1
		icmpMsg[6] = 0
		icmpMsg[7] = 2

		checksum := checksum(icmpMsg)
		icmpMsg[2] = byte(checksum >> 8)
		icmpMsg[3] = byte(checksum)

		// msg := icmp.Message {
		// 	Type: ipv4.ICMPTypeEcho, Code: 0,
		// 	Body: &icmp.Echo{
		// 		ID: os.Getpid() & 0xffff, Seq: 1,
		// 		Data: []byte(""),
		// 	},
		// }
		// icmpMsg, err := msg.Marshal(nil)
		// if err != nil {
		//     fmt.Println("Error creating icmpMsg: ", err)
		//     return
		// }

		// Send ICMP packet
		start := time.Now()

		_, err = conn.WriteTo(icmpMsg, ipAddr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		buff := make([]byte, 512)
		err = conn.SetReadDeadline(time.Now().Add(250 * time.Millisecond))
		if err != nil {
			fmt.Println("set read deadline error: ", err)
			return
		}

		n, addr, err := conn.ReadFrom(buff)
		if err != nil {
			// fmt.Println("Read error: ", err)
			fmt.Println("*\t*\t*")
			continue
		}

		duration := time.Since(start)

		rm, err := icmp.ParseMessage(ProtocolICMP, buff[:n])
		if err != nil {
			fmt.Println(err)
			return
		}

		switch rm.Type {

		case ipv4.ICMPTypeEchoReply:
			fmt.Println(ipAddr, ttl, duration)
			os.Exit(1)

		case ipv4.ICMPTypeTimeExceeded:
			fmt.Println(&net.IPAddr{IP: addr.(*net.IPAddr).IP}, ttl, duration)

		default:
			fmt.Println("got %+v from %v; want echo reply", rm, addr)
		}
	}
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
