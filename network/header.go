package network

import (
	"encoding/binary"
)

type header struct {
	Checksum   uint32
	Ack        uint8
	AckReturn  uint8
	PacketType packettype_t
	Reserved   uint8
}

type checksumCalculator struct {
	checksum uint32
	term     int
}

func newChecksumCalculator() *checksumCalculator {
	return &checksumCalculator{0x1234567, 1}
}

func (c *checksumCalculator) Write(data []byte) (n int, err error) {
	var skip = 0
	// if term is the initial value, we skip the checksum of the header
	if c.term == 1 {
		skip = 4
	}
	for _, d := range data[skip:] {
		c.checksum += uint32(d) * uint32(c.term)
		c.term += 1
	}
	return len(data), nil
}

func (c *checksumCalculator) setChecksum(h *header, data any) {
	err := binary.Write(c, binary.LittleEndian, h)
	if err != nil {
		panic(err)
	}
	err = binary.Write(c, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}
	h.Checksum = c.checksum
}

func makeHeader(packetType packettype_t, packetData any) header {
	h := header{0, 0, 0, packetType, 0}
	c := newChecksumCalculator()
	c.setChecksum(&h, packetData)
	return h
}
