package main

import (
	"log"
	"net"
)

const PORT = 55001

func main() {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: PORT})
	if err != nil {
		log.Fatalf("Issue connecting to UDP socket: %s\n", err)
	}
	defer conn.Close()
	log.Printf("Connected to Port %d\n", PORT)

	clients := make(map[string]*net.UDPAddr)
	buf := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("An error occurred while attempting to read client message. Skipping message.")
			continue
		}

		key := clientAddr.String()
		message := string(buf[:n])
		if _, exists := clients[key]; !exists {
			clients[key] = clientAddr
			log.Printf("New client joined: %s\n", clientAddr)
		}

		for addrStr, c := range clients {
			if addrStr == key {
				continue
			}
			_, err := conn.WriteToUDP([]byte(message), c)
			if err != nil {
				log.Printf("An error occurred while sending a message to client %s\n", addrStr)
				delete(clients, addrStr)
			}
		}
	}
}
