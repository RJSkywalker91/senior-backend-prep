package main

import (
	common "chat-service/shared"
	"encoding/binary"
	"log"
)

func SuccessPath() {
	log.Println("-----------GREEN PATH TEST-----------")
	log.Println("Creating a message packet...")
	message := common.MessagePacket{Name: "RJSkywalker", Message: "This is a test!"}
	log.Println("Marshalling message packet...")
	packet, _ := common.Marshal(&message)
	log.Println("Pretending byte[] being sent across the wire...")
	log.Println("Unmarshalling message packet")
	unmarshalled, err := common.UnMarshal(packet)
	if err != nil {
		log.Fatalf("\nSomething went wrong unmarshalling the message packet...\n%s\n", err)
	}
	log.Printf("Unmarshalled values...\nName: %s\nMessage: %s", unmarshalled.Name, unmarshalled.Message)
	log.Println("-----------GREEN PATH TEST FINISHED-----------")
}

func NameErrorTests() {
	log.Println("-----------NAME ERROR TESTS-----------")
	log.Println("Name < 3 characters")
	p := common.MessagePacket{Name: "A", Message: ""}
	err := p.ValidateMessagePacket()
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}
	log.Println("Name > 50 characters")
	p = common.MessagePacket{Name: "Supercalafragelisticexpealidocious01234567899876543210", Message: ""}
	err = p.ValidateMessagePacket()
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}
	log.Println("-----------NAME ERROR TESTS FINISHED-------")
}

func MarshalErrorTest() {
	log.Println("-----------MARSHAL ERROR TESTS-----------")
	log.Println("nil MessagePacket")
	_, err := common.Marshal(nil)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}
	log.Println("Invalid MessagePacket")
	message := common.MessagePacket{Name: "A", Message: "This is a test!"}
	_, err = common.Marshal(&message)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}
	log.Println("-----------MARSHAL ERROR TESTS FINISHED-------")
}

func UnMarshalErrorTests() {
	log.Println("-----------UNMARSHAL ERROR TESTS-----------")
	log.Println("Short Buffer: Name Length")
	buffer := make([]byte, 3)
	_, err := common.UnMarshal(buffer)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}

	log.Println("Short Buffer: Name")
	buffer = make([]byte, 7)
	binary.BigEndian.PutUint32(buffer[0:4], uint32(4))
	copy(buffer[4:7], "ABC")
	_, err = common.UnMarshal(buffer)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}

	log.Println("Short Buffer: Message Length")
	buffer = make([]byte, 11)
	binary.BigEndian.PutUint32(buffer[0:4], uint32(4))
	copy(buffer[4:8], "TEST")
	_, err = common.UnMarshal(buffer)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}

	log.Println("Short Buffer: Message")
	buffer = make([]byte, 15)
	binary.BigEndian.PutUint32(buffer[0:4], uint32(4))
	copy(buffer[4:8], "TEST")
	binary.BigEndian.PutUint32(buffer[8:12], uint32(4))
	_, err = common.UnMarshal(buffer)
	if err == nil {
		log.Println("TEST FAILED.")
	} else {
		log.Printf("TEST SUCCEEDED with err: \"%s\"\n", err)
	}
	log.Println("-----------UNMARSHAL ERROR TESTS FINISHED-----------")
}

func main() {
	NameErrorTests()
	MarshalErrorTest()
	UnMarshalErrorTests()
	SuccessPath()
}
