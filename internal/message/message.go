package message

import (
	"encoding/binary"
)

const (
	messageSigningDomain = "\xffsolana offchain"
)

type OffchainMessageOpts struct {
	MessageBody []byte
	Version     uint8
}

func CreateOffchainMessageWithPreamble(opts *OffchainMessageOpts) []byte {
	// Signing domain (16 bytes) + Version (1 byte) + Format (1 byte) + Len (2 bytes)
	preamble := make([]byte, 0, 20+len(opts.MessageBody))

	// Signing domain
	preamble = append(preamble, []byte(messageSigningDomain)...)

	// Header version
	preamble = append(preamble, byte(opts.Version))

	// Message format
	preamble = append(preamble, 0x00)

	// Message length
	lenBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(lenBytes, uint16(len(opts.MessageBody)))
	preamble = append(preamble, lenBytes...)

	// Message body
	preamble = append(preamble, opts.MessageBody...)

	return preamble
}
