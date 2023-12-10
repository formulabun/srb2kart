package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func scanFiles(response filesNeeded) ([]file, error) {
	files := make([]file, 0, response.Num)
	fileScanner := bufio.NewScanner(bytes.NewBuffer(response.Files[:]))
	fileScanner.Split(scanFile)
	for i := 0; i < int(response.Num); i++ {
		if !fileScanner.Scan() {
			return files, fmt.Errorf("Could not read the next file needed: %s", fileScanner.Err())
		}
		f, err := fileTokenToFile(fileScanner.Bytes())
		if err != nil {
			return files, fmt.Errorf("Could not read files needed: %s", err)
		}
		files = append(files, f)
	}

	return files, nil
}

func scanFile(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// see file struct. WadName is null terminated if shorter than MAX_WADPATH (512)
	preambleLen := 5
	postamleLen := 1 + 16
	if len(data) <= 9 { // too small
		return 0, nil, nil
	}
	if data[0] == 0 {
		return 0, []byte{}, bufio.ErrFinalToken
	}
	// start searching for null termination after status and size
	term := bytes.Index(data[preambleLen:], []byte{0x0}) + preambleLen
	if term < 0 {
		return 0, nil, nil
	}
	if term == preambleLen {
		return 0, nil, errors.New("Could not read filename")
	}
	token = make([]byte, term+postamleLen) // until null term + 0b0 + checksum(16)
	copy(token, data)
	return term + postamleLen, token, nil
}

func fileTokenToFile(data []byte) (file, error) {
	checksum := [16]byte{}
	copy(checksum[:], data[len(data)-16:])
	return file{
		uint8(data[0]),
		binary.LittleEndian.Uint32(data[1:5]),
		string(data[5 : len(data)-16-1]), // -(checksum size + 1 null byte)
		checksum,
	}, nil
}
