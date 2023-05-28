package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func AskInfo(address string) (ServerInfo, []PlayerInfo, error) {

	conn, err := SendPacket(address, pt_askinfo, askInfo{0x1, 0x21f})
	if err != nil {
		return ServerInfo{}, nil, fmt.Errorf("Could not ask for server info: %s", err)
	}
	defer conn.Close()

  // read server info
	var server = ServerInfo{}
	var buffer = make([]byte, binary.Size(server))
  head, err := ReadHeader(conn, buffer)
  if err != nil {
    return server, nil, fmt.Errorf("Could not get server info response: %s", err)
  }

  buf := bytes.NewBuffer(buffer)
  binary.Read(buf, binary.LittleEndian, &server)

  if head.packettype != pt_serverinfo {
    panic("Expected server info packet to come first")
  }

  // read players
	var players = make([]PlayerInfo, server.MaxPlayer)
  buffer = make([]byte, binary.Size(players))
  head, err = ReadHeader(conn, buffer)
  if err != nil {
    return server, players, fmt.Errorf("Could not get server info response: %s", err)
  }

  buf = bytes.NewBuffer(buffer)
  binary.Read(buf, binary.LittleEndian, &players)

  if head.packettype != pt_playerinfo {
    panic("Expected to receive packet info")
  }

  return server, players, nil
}

func TellFilesNeeded() (any, error) {
  return nil, nil
}
