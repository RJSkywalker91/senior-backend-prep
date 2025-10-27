package common

import (
	"encoding/binary"
	"errors"
)

type MessagePacket struct {
	Name    string
	Message string
}

func (p *MessagePacket) ValidateMessagePacket() error {
	if len(p.Name) < 3 || len(p.Name) > 50 {
		return errors.New("name must be between 3 and 50 characters")
	}

	if len(p.Message) > 200 {
		return errors.New("message must be less than 200 characters")
	}

	return nil
}

func Marshal(p *MessagePacket) ([]byte, error) {
	if p == nil || p.ValidateMessagePacket() != nil {
		return nil, errors.New("unable to proceed with invalid MessagePacket")
	}

	nameLength := len(p.Name)
	messageLength := len(p.Message)
	byteSize := 8 + nameLength + messageLength
	buffer := make([]byte, byteSize)

	binary.BigEndian.PutUint32(buffer[0:4], uint32(nameLength))
	copy(buffer[4:4+nameLength], p.Name)

	binary.BigEndian.PutUint32(buffer[4+nameLength:8+nameLength], uint32(messageLength))
	copy(buffer[8+nameLength:], p.Message)

	return buffer, nil
}

func UnMarshal(data []byte) (MessagePacket, error) {
	var packet MessagePacket
	position := 0

	// Read first 4 bytes to see how long Name is (n)
	if len(data)-position < 4 {
		return packet, errors.New("short buffer: name length")
	}
	nameLength := int(binary.BigEndian.Uint32(data[position:]))
	position += 4

	// Read n bytes and store in packet.Name
	if len(data)-position < nameLength {
		return packet, errors.New("short buffer: name")
	}
	packet.Name = string(data[position : position+nameLength])
	position += nameLength

	// Read next 4 bytes to see how long Message is (m)
	if len(data)-position < 4 {
		return packet, errors.New("short buffer: message length")
	}
	messageLength := int(binary.BigEndian.Uint32(data[position:]))
	position += 4

	// Read m bytes and store in packet.Message
	if len(data)-position < messageLength {
		return packet, errors.New("short buffer: message")
	}
	packet.Message = string(data[position : position+messageLength])
	position += messageLength

	return packet, nil
}
