package network

import (
	"testing"
)

func TestHeaderChecksum(t *testing.T) {
  // endianness is manually checked in this test
	header := makeHeader(pt_askinfo, [5]byte{0x1, 0x1f, 0x2, 0x0, 0x0})
  expected := uint32(0x01234658)
  t.Logf("header in hex is %x", header)
  if header.checksum != expected {
    t.Fatalf("Expected a checksum of 0x%x, but got 0x%x", expected, header.checksum)
  }
}

func TestHeaderAskInfo(t *testing.T) {
  // endianness is manually checked in this test
	header := makeHeader(pt_askinfo, askInfo{0x1, 0x21f})
  expected := uint32(0x01234658)
  t.Logf("header in hex is %x", header)
  if header.checksum != expected {
    t.Fatalf("Expected a checksum of 0x%x, but got 0x%x", expected, header.checksum)
  }
}
