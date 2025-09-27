package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create client socket
	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 5001,
	}
	client, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Issue creating client socket: %s\n", err)
	}
	defer client.Close()

	fmt.Println("Listening for broadcasts from time server...")
	data := make([]byte, 1024)
	for {
		n, err := client.Read(data)
		if err != nil {
			fmt.Print("An error occurred when attempting to read")
			continue
		}
		if n > 0 {
			fmt.Printf("Server Time: %s\n", string(data[:n]))
		}
	}
}
