package main

import (
	"fmt"
	"net"
)

const maxPackageLength = 60000
const udpAddress = ":8030"

var count = 0

func onRecv(addr *net.UDPAddr, msg string) {
	fmt.Println("Received ", msg, " from ", addr, ". Total length is ", count)
}

func startServer(address string) {
	serverAddr, err := net.ResolveUDPAddr("udp", address)
	ExitOnError(err)

	/* Now listen at selected port */
	serverConn, err := net.ListenUDP("udp", serverAddr)
	ExitOnError(err)
	defer serverConn.Close()

	buf := make([]byte, maxPackageLength)

	fmt.Println("UDP server start listening on ", serverAddr, " ...")

	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		count += n

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			go onRecv(addr, string(buf[0:n]))
		}
	}
}

func main() {
	startServer(udpAddress)
}
