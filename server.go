package main

import (
	"fmt"
	"net"
)

const maxPackageLength = 60000
const udpAddress = ":8030"

func main() {

	serverAddr, err := net.ResolveUDPAddr("udp", udpAddress)
	ExitOnError(err)

	/* Now listen at selected port */
	serverConn, err := net.ListenUDP("udp", serverAddr)
	ExitOnError(err)
	defer serverConn.Close()

	buf := make([]byte, maxPackageLength)

	count := 0

	fmt.Println("UDP server start listening on ", serverAddr, " ...")

	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		count += n
		fmt.Println("Received ", string(buf[0:n]), " from ", addr, ". Total length is ", count)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}

}
