package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func AskInfo(address string) (ServerInfo, []PlayerInfo, error) {
	conn, err := SendPacket(address, pt_askinfo, askInfo{0x0, 0x0})
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

	if head.PacketType != pt_serverinfo {
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

	if head.PacketType != pt_playerinfo {
		panic("Expected to receive packet info")
	}

	return server, players, nil
}

func TellAllFilesNeeded(address string) ([]string, error) {
	conn, err := OpenConnection(address)
	if err != nil {
		fmt.Printf("Could not ask for needed files: %s", err)
	}
	defer conn.Close()

	filesNeeded, files, err := tellFilesNeeded(conn, 0)
	if err != nil {
		return nil, fmt.Errorf("Could not request files: %s", err)
	}
	result := make([]string, 0, filesNeeded.Num)
	for _, f := range files {
		result = append(result, f.WadName)
	}
	for filesNeeded.More > 0 {
		filesNeeded, files, err = tellFilesNeeded(conn, len(result))
		if err != nil {
			return nil, fmt.Errorf("Could not request files: %s", err)
		}
		for _, f := range files {
			result = append(result, f.WadName)
		}
	}

	return result, nil
}

func tellFilesNeeded(conn *net.UDPConn, from int) (FilesNeeded, []File, error) {
	fmtErr := func(err error) (FilesNeeded, []File, error) {
		return FilesNeeded{}, nil, fmt.Errorf("Could get needed files: %s", err)
	}

	err := SendPacketOnConnection(conn, pt_tellfilesneeded, FilesNeededNum(from))
	if err != nil {
		return fmtErr(err)
	}

	var response FilesNeeded
	buffer := make([]byte, binary.Size(response))
	head, err := ReadHeader(conn, buffer)
	if err != nil {
		return fmtErr(err)
	}
	if head.PacketType != pt_morefilesneeded {
		return FilesNeeded{}, nil, fmt.Errorf("Unexpected packet response: %x instead of %x", head.PacketType, pt_morefilesneeded)
	}

	buff := bytes.NewBuffer(buffer)
	err = binary.Read(buff, binary.LittleEndian, &response)
	if err != nil {
		return fmtErr(err)
	}

	files := make([]File, 0, response.Num)
	fileScanner := bufio.NewScanner(bytes.NewBuffer(response.Files[:]))
	fileScanner.Split(ScanFile)
  for i := 0 ; i < int(response.Num); i++ {
    if !fileScanner.Scan() {
      return response, files, fmt.Errorf("Could not read the next file needed: %s", fileScanner.Err())
    }
		f, err := fileTokenToFile(fileScanner.Bytes())
		if err != nil {
			return response, files, fmt.Errorf("Could not read files needed: %s", err)
		}
		files = append(files, f)
	}

	return response, files, nil
}
