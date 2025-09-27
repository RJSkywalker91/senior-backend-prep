package main

import (
	"fmt"
	"net"
)

const NETWORK = "tcp"
const HOST = "127.0.0.1:80"

func main() {
	// Create client socket
	client, err := net.Dial(NETWORK, HOST)
	if err != nil {
		fmt.Printf("Issue creating client socket: %s", err)
		return
	}

	// Send data to the server
	message := "Hello World!"
	client.Write([]byte(message))

	// Get response from the server
	response := make([]byte, 1024)
	client.Read(response)
	fmt.Printf("Response from server: %s", string(response))

	// Close socket
	client.Close()
}
