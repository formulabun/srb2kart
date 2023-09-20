package network

import "fmt"

type Checksum [16]byte

type File struct {
	Filename string
	Checksum Checksum
}

func (c Checksum) String() string {
	return fmt.Sprintf("%1x", [16]byte(c))
}
