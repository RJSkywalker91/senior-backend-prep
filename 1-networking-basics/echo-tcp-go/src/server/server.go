package main

import (
	"fmt"
	"net"
)

const NETWORK = "tcp"
const HOST = "127.0.0.1:80"

func main() {
	// Create socket + listen
	server, err := net.Listen(NETWORK, HOST)
	if err != nil {
		fmt.Printf("\nIssue starting server: %s", err)
		return
	}
	fmt.Print("\nServer listening on Port 80")

	// Accept a connection
	client, connErr := server.Accept()
	if connErr != nil {
		fmt.Print("\nError connecting client to server... exiting application")
		server.Close()
		return
	}

	// Echo received data back to connection
	data := make([]byte, 1024)
	_, readErr := client.Read(data)
	if readErr != nil {
		fmt.Print("\nError reading client message... exiting application")
		server.Close()
		return
	}
	fmt.Printf("\nMessage received from client: %s", string(data))
	client.Write(data)

	// Close sockets
	client.Close()
	server.Close()
	fmt.Print("\nClient and Server sockets closed")
}
