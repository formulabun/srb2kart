package wad

import (
	"encoding/binary"
	"errors"
	"io"
)

type T int

func ReadWad(reader io.ReadSeeker) (result Wad, err error) {
	err = binary.Read(reader, binary.LittleEndian, &result.Header)
	if err != nil {
		return
	}
	numwads := result.NumWads
	_, err = reader.Seek(int64(result.Header.InfoTableOffset), io.SeekStart)
	if err != nil {
		return
	}
	result.Directory = make([]*FileLump, numwads)

	for i, _ := range result.Directory {
		var lumpInfo FileLump
		err = binary.Read(reader, binary.LittleEndian, &lumpInfo)
		if err != nil {
			return
		}
		result.Directory[i] = &lumpInfo
	}
	return
}

func ReadLump(reader io.ReadSeeker, lump *FileLump) ([]byte, error) {
	result := make([]byte, lump.Size)
	_, err := reader.Seek(int64(lump.FilePos), io.SeekStart)
	if err != nil {
		return result, err
	}
	n, err := reader.Read(result)
	if err != nil {
		return result, err
	}
	if n != int(lump.Size) {
		return result, errors.New("Could not read all bytes of the lump")
	}
	return result, nil
}
