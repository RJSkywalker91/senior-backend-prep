package main

import (
	common "chat-service/shared"
	"log"
	"net"
)

const PORT = 55001
const SERVER_ADDRESS_STR = ":55001"

type msgEvent struct {
	addr  *net.UDPAddr
	bytes []byte
}

func runChat(conn *net.UDPConn, events <-chan msgEvent) {
	clients := make(map[string]*net.UDPAddr)
	broadcast := func(e msgEvent) {
		for k, c := range clients {
			if k != e.addr.String() {
				log.Println("Sending message to client")
				_, _ = conn.WriteToUDP(e.bytes, c)
			}
		}
	}

	for event := range events {
		key := event.addr.String()
		_, ok := clients[key]
		if !ok {
			log.Println("New client joined the chat")
			clients[key] = event.addr
		}
		messagePacket, err := common.UnMarshal(event.bytes)
		if err != nil {
			log.Printf("\nSomething went wrong unmarshalling the message packet for addr %s\n%s\n", event.addr, err)
		}
		if err := messagePacket.ValidateMessagePacket(); err != nil {
			log.Printf("Message Packet is invalid. Unable to broadcast.")
			continue
		}
		broadcast(event)
	}
}

func main() {
	// Bind UDP
	serverAddr, err := net.ResolveUDPAddr("udp", SERVER_ADDRESS_STR)
	if err != nil {
		log.Fatalf("Unable to resolve address for:")
	}

	// Connect and listen on server address
	conn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		log.Fatalf("Issue connecting to UDP socket: %s\n", err)
	}
	defer conn.Close()
	log.Printf("UDP custom server listening on :%d\n", PORT)

	// Startup Chat Service
	msgEvents := make(chan msgEvent, 1024)
	go runChat(conn, msgEvents)

	buf := make([]byte, 2048)

	// Event Loop for reading
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("An error occurred while attempting to read client message. Skipping message.")
			continue
		}
		log.Printf("Sending message")
		msgEvents <- msgEvent{addr: addr, bytes: buf[:n]}
	}
}
