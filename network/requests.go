package network

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

func AskInfo(address string) (ServerInfo, []PlayerInfo, error) {
	conn, err := sendPacket(address, pt_askinfo, askInfo{0x0, 0x0})
	if err != nil {
		return ServerInfo{}, nil, fmt.Errorf("Could not ask for server info: %s", err)
	}
	defer conn.Close()

	// read servers
	var server = ServerInfo{}
	var buffer = make([]byte, binary.Size(server))
	head, err := readHeader(conn, buffer)
	if err != nil {
		return server, nil, fmt.Errorf("Could not get server info response: %s", err)
	}

	if head.PacketType != pt_serverinfo {
		return server, []PlayerInfo{}, errors.New("Received info packets out of sync")
	}

	buf := bytes.NewBuffer(buffer)
	binary.Read(buf, binary.LittleEndian, &server)

	// read players
	var players = make([]PlayerInfo, server.MaxPlayer)
	buffer = make([]byte, binary.Size(players))
	head, err = readHeader(conn, buffer)
	if err != nil {
		return server, players, fmt.Errorf("Could not get server info response: %s", err)
	}

	if head.PacketType != pt_playerinfo {
		return server, players, errors.New("Received info packets out of sync")
	}

	buf = bytes.NewBuffer(buffer)
	binary.Read(buf, binary.LittleEndian, &players)

	return server, players, nil
}

func TellAllFilesNeeded(address string) ([]File, error) {
	conn, err := openConnection(address)
	if err != nil {
		fmt.Printf("Could not ask for needed files: %s", err)
		return []File{}, err
	}
	defer conn.Close()

	filesNeeded, files, err := tellFilesNeeded(conn, 0)
	if err != nil {
		return nil, fmt.Errorf("Could not request files: %s", err)
	}
	result := make([]File, 0, filesNeeded.Num)
	for _, f := range files {
		result = append(result, File{f.WadName, f.Md5Sum})
	}
	for filesNeeded.More > 0 {
		filesNeeded, files, err = tellFilesNeeded(conn, len(result))
		if err != nil {
			return nil, fmt.Errorf("Could not request files: %s", err)
		}
		for _, f := range files {
			result = append(result, File{f.WadName, f.Md5Sum})
		}
	}

	return result, nil
}

func tellFilesNeeded(conn *net.UDPConn, from int) (filesNeeded, []file, error) {
	fmtErr := func(err error) (filesNeeded, []file, error) {
		return filesNeeded{}, nil, fmt.Errorf("Could get needed files: %s", err)
	}

	err := sendPacketOnConnection(conn, pt_tellfilesneeded, filesNeededNum(from))
	if err != nil {
		return fmtErr(err)
	}

	var response filesNeeded
	buffer := make([]byte, binary.Size(response))
	head, err := readHeader(conn, buffer)
	if err != nil {
		return fmtErr(err)
	}
	if head.PacketType != pt_morefilesneeded {
		return filesNeeded{}, nil, fmt.Errorf("Unexpected packet response: %x instead of %x", head.PacketType, pt_morefilesneeded)
	}

	buff := bytes.NewBuffer(buffer)
	err = binary.Read(buff, binary.LittleEndian, &response)
	if err != nil {
		return fmtErr(err)
	}

	files, err := scanFiles(response)

	return response, files, err
}
