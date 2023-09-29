package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"

	"golang.org/x/net/ipv6"
)

var port = flag.Int("port", 514, "UDP port number")
var address = flag.String("address", "ff05::514", "IPv6 multicast destination")

func main() {
	flag.Parse()
	listen := net.JoinHostPort("::", strconv.Itoa(*port))
	conn, err := net.ListenPacket("udp6", listen)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	p := ipv6.NewPacketConn(conn)
	group := net.UDPAddr{IP: net.ParseIP(*address)}
	if err := p.JoinGroup(nil, &group); err != nil {
		panic(err)
	}
	if err := p.SetMulticastLoopback(true); err != nil {
		panic(err)
	}
	for {
		buffer := make([]byte, 9000)
		n, _, err := conn.ReadFrom(buffer)
		if err != nil {
			panic(err)
		}
		s := string(buffer[:n])
		s = strings.TrimSpace(s)
		fmt.Println(s)
	}
}
