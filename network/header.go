package network

import (
	"encoding/binary"
)

type header struct {
	checksum   uint32
	ack        uint8
	ackreturn  uint8
	packettype uint8
	reserved   uint8
}

type checksumCalculator struct {
	checksum uint32
	offset   uint32
}

func newChecksumCalculator() *checksumCalculator {
	return &checksumCalculator{0x1234567, 1}
}

func (c *checksumCalculator) Write(data []byte) (n int, err error) {
	for _, d := range data {
		c.checksum += uint32(d) * c.offset
		c.offset += 1
	}
	return len(data), nil
}

func (c *checksumCalculator) setChecksum(h *header, data any) {
	err := binary.Write(c, binary.LittleEndian, h.ack)
	if err != nil {
		panic(err)
	}
	err = binary.Write(c, binary.LittleEndian, h.ackreturn)
	if err != nil {
		panic(err)
	}
	err = binary.Write(c, binary.LittleEndian, h.packettype)
	if err != nil {
		panic(err)
	}
	err = binary.Write(c, binary.LittleEndian, h.reserved)
	if err != nil {
		panic(err)
	}
	err = binary.Write(c, binary.LittleEndian, data)
	if err != nil {
		panic(err)
	}
	h.checksum = c.checksum
}

func makeHeader(packetType uint8, packetData any) header {
	h := header{0, 0, 0, packetType, 0}
	c := newChecksumCalculator()
	c.setChecksum(&h, packetData)
	return h
}
