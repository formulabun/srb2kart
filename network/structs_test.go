package network

import "testing"

func TestChecksumString(t *testing.T) {
	c := Checksum{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf}
	s := c.String()
	if s != "000102030405060708090a0b0c0d0e0f" {
		t.Fatalf("Checksum{0x0123456789abcdef}.String() = '%s' but expected '000102030405060708090a0b0c0d0e0f'", s)
	}
}
