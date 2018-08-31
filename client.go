package main

import (
	"fmt"
	"net"
	"sync"
)

const serverAddress = "127.0.0.1:8030"

func send(addr string, buffer []byte) {
	connect, err := net.Dial("udp", addr)
	defer connect.Close()

	ExitOnError(err)

	_, err = connect.Write(buffer)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	text := []string{"hello world", "12345", "67890"}

	var wg sync.WaitGroup

	fmt.Println("start sending udp message")
	for _, msg := range text {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			send(serverAddress, []byte(msg))
		}(msg)
	}

	wg.Wait()

	fmt.Println("udp sending completed!")
}
