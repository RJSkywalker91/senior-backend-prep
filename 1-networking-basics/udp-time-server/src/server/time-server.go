package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr := &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 5001,
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Issue connecting to UDP socket: %s\n", err)
	}
	fmt.Println("Connected to broadcast IP on Port 5001")
	defer conn.Close()

	// Write time to clients every 5 seconds
	for {
		currTime := time.Now()
		conn.Write([]byte(currTime.String()))
		fmt.Println("Time broadcasted")
		time.Sleep(5 * time.Second)
	}
}
