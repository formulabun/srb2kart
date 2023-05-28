package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

func SendPacket(address string, packetType uint8, packetData any) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, fmt.Errorf("Could not resolve address %s: %s", address, err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return nil, fmt.Errorf("Could not open connection to %s: %s", address, err)
	}

	conn.SetReadBuffer(2048)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	err = SendPacketOnConnection(conn, packetType, packetData)

	return conn, err
}

func SendPacketOnConnection(conn *net.UDPConn, packetType uint8, packetData any) error {
	packet := makeHeader(packetType, packetData)
	var err error
	var buff bytes.Buffer
	err = binary.Write(&buff, binary.LittleEndian, packet.checksum)
	if err != nil {
		return err
	}
	err = binary.Write(&buff, binary.LittleEndian, packet.ack)
	if err != nil {
		return err
	}
	err = binary.Write(&buff, binary.LittleEndian, packet.ackreturn)
	if err != nil {
		return err
	}
	err = binary.Write(&buff, binary.LittleEndian, packet.packettype)
	if err != nil {
		return err
	}
	err = binary.Write(&buff, binary.LittleEndian, packet.reserved)
	if err != nil {
		return err
	}
	err = binary.Write(&buff, binary.LittleEndian, packetData)
	if err != nil {
		return err
	}
	_, err = io.Copy(conn, &buff)
	return err
}

func ReadHeader(conn *net.UDPConn, data []byte) (h header, err error) {
  var bs = make([]byte, 8 + len(data)) // enough for the header
  n, err := conn.Read(bs)
  if err != nil {
    return h, fmt.Errorf("Could not read resonse header after %d bytes: %s", n, err)
  }

  buf := bytes.NewBuffer(bs)
	err = binary.Read(buf, binary.LittleEndian, &h.checksum)
	if err != nil {
		return h, fmt.Errorf("Could not read checksum from connection: %s", err)
	}
	err = binary.Read(buf, binary.LittleEndian, &h.ack)
	if err != nil {
		return h, fmt.Errorf("Could not read ack from connection: %s", err)
	}
	err = binary.Read(buf, binary.LittleEndian, &h.ackreturn)
	if err != nil {
		return h, fmt.Errorf("Could not read ack return from connection: %s", err)
	}
	err = binary.Read(buf, binary.LittleEndian, &h.packettype)
	if err != nil {
		return h, fmt.Errorf("Could not read packet type from connection: %s", err)
	}
	err = binary.Read(buf, binary.LittleEndian, &h.reserved)
	if err != nil {
		return h, fmt.Errorf("Could not read reserved byte from connection: %s", err)
	}
  buf.Read(data)
	return
}
