package common

import (
	"encoding/binary"
	"errors"
	"hash/crc32"
)

// message type
const (
	PING = uint8(iota)
	PONG
	CREATE_TOPIC
	DELETE_TOPIC
	SEND_MESSAGE
	GET_MESSAGE
)

const (
	ERROR_CRC       = "crc error"
	TOPIC_EXIST     = "topic exist"
	TOPIC_NOT_EXIST = "topic not exist"
)

type Header struct {
	Type   uint8
	UserID uint8
}

type Message struct {
	Length uint32
	CRC    uint32
	Header Header
	Body   []byte
}

// marshal message to bytes
func (m *Message) Marshal() []byte {
	// create binary message
	msg := make([]byte, 8+2+len(m.Body))
	// put data to msg
	binary.BigEndian.PutUint32(msg[0:4], m.Length)
	msg[8] = m.Header.Type
	msg[9] = m.Header.UserID
	copy(msg[10:], m.Body)
	// calculation CRC value of msg

	m.CRC = crc32.ChecksumIEEE(msg)
	// put CRC value to msg
	binary.BigEndian.PutUint32(msg[4:8], m.CRC)
	return msg
}

// umarshal message from bytes
func Unmarshal(msg []byte) (*Message, error) {
	message := &Message{}
	message.Length = binary.BigEndian.Uint32(msg[0:4])
	message.CRC = binary.BigEndian.Uint32(msg[4:8])
	message.Header.Type = msg[8]
	message.Header.UserID = msg[9]
	message.Body = msg[10:]
	// check CRC
	if message.CRC != crc32.ChecksumIEEE(msg) {
		return nil, errors.New(ERROR_CRC)
	}
	return message, nil
}
