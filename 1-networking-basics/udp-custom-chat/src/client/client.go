package main

import (
	"bufio"
	common "chat-service/shared"
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
		msg := sc.Text()
		if len(msg) == 0 {
			continue
		}
		p := common.MessagePacket{Name: conn.LocalAddr().String(), Message: msg}
		packetBytes, err := common.Marshal(&p)
		if err != nil {
			log.Fatalf("an error occurred during marshalling of message: %s", err)
		}
		if _, err := conn.Write(packetBytes); err != nil {
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
		log.Println("Data read from buffer. Unmarshalling...")
		messagePacket, err := common.UnMarshal(buf[:n])
		if err != nil {
			log.Printf("\nSomething went wrong unmarshalling the message packet\n%s\n", err)
			continue
		}
		msg := messagePacket.Name + ": " + messagePacket.Message
		fmt.Printf("<< %s\n", msg)
	}
}
