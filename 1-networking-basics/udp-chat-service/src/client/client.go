package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const PORT = 55001

func main() {
	serverAdd := &net.UDPAddr{Port: PORT}
	conn, err := net.DialUDP("udp", nil, serverAdd)
	if err != nil {
		log.Fatalf("Issue connecting client to server socket: %s\n", err)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", serverAdd)

	go readRoutine(conn)

	sc := bufio.NewScanner(os.Stdin)
	for {
		log.Print(">> ")
		if !sc.Scan() {
			break
		}
		line := sc.Bytes()
		if len(line) == 0 {
			continue
		}
		if _, err := conn.Write(line); err != nil {
			log.Printf("write error: %v\n", err)
		}
	}
	if err := sc.Err(); err != nil {
		log.Printf("stdin error: %v\n", err)
	}

	fmt.Println("Shutting down.")
}

func readRoutine(conn *net.UDPConn) {
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf) // reads only from the dialed peer
		if err != nil {
			log.Printf("read error: %v\n", err)
			return
		}
		fmt.Printf("<< %s\n", string(buf[:n]))
	}
}
