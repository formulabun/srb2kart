package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func ScanFile(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) <= 9 { // too small
		return 0, nil, nil
	}
	if data[0] == 0 {
		return 0, []byte{}, bufio.ErrFinalToken
	}
	term := bytes.Index(data[5:], []byte{0x0}) + 6
	if term < 0 {
		return 0, nil, nil
	}
	token = make([]byte, term+16)
	copy(token, data)
	return term + 16, token, nil
}

func fileTokenToFile(data []byte) (File, error) {
	checksum := [16]byte{}
	copy(checksum[:], data[len(data)-16:])
	return File{
		uint8(data[0]),
		binary.LittleEndian.Uint32(data[1:5]),
		string(data[5 : len(data)-16]),
		checksum,
	}, nil
}
