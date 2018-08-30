package main

import (
	"fmt"
	"net"
)

func main() {
	connect, err := net.Dial("udp", "127.0.0.1:8030")
	ExitOnError(err)

	defer connect.Close()

	text := []string{"hello world", "12345", "67890"}

	fmt.Print("start sending udp message")
	for _, msg := range text {
		buf := []byte(msg)
		_, err = connect.Write(buf)
		if err != nil {
			fmt.Println(msg, err)
		}
	}

	fmt.Println("udp sending completed!")
}
