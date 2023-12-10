package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
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
	if len(data) <= 9 { // too small
		return 0, nil, nil
	}
	if data[0] == 0 {
		return 0, []byte{}, bufio.ErrFinalToken
	}
	term := bytes.Index(data[6:], []byte{0x0}) + 6
	if term < 0 {
		return 0, nil, nil
	}
	token = make([]byte, term+16)
	copy(token, data)
	return term + 16, token, nil
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
