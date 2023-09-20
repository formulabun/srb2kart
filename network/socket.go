package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

func openConnection(address string) (*net.UDPConn, error) {
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

	return conn, nil
}

func sendPacket(address string, packetType packettype_t, packetData any) (*net.UDPConn, error) {
	conn, err := openConnection(address)
	if err != nil {
		return nil, fmt.Errorf("Could not open socket for sending a packet: %s", err)
	}

	err = sendPacketOnConnection(conn, packetType, packetData)

	return conn, err
}

func sendPacketOnConnection(conn *net.UDPConn, packetType packettype_t, packetData any) error {
	packet := makeHeader(packetType, packetData)
	var err error
	var buff bytes.Buffer
	err = binary.Write(&buff, binary.LittleEndian, packet)
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

func readHeader(conn *net.UDPConn, data []byte) (h header, err error) {
	var bs = make([]byte, 8+len(data)) // enough for the header
	n, err := conn.Read(bs)
	if err != nil {
		return h, fmt.Errorf("Could not read resonse header after %d bytes: %s", n, err)
	}

	buf := bytes.NewBuffer(bs)
	err = binary.Read(buf, binary.LittleEndian, &h)
	if err != nil {
		return h, fmt.Errorf("Could not read header from connection: %s", err)
	}
	buf.Read(data)
	return
}
