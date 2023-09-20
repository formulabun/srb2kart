package wad

import (
	"bytes"
	"encoding/binary"
	"io"
)

type T int

func ReadWad(reader io.Reader) (result Wad, err error) {
	result.data, err = io.ReadAll(reader)
	if err != nil {
		return
	}

	headerSize := binary.Size(&result.Header)
	headerBuff := bytes.NewBuffer(result.data[0:headerSize])
	err = binary.Read(headerBuff, binary.LittleEndian, &result.Header)
	if err != nil {
		return
	}
	if err = validateHeader(&result.Header); err != nil {
		return result, err
	}

	numwads := result.NumWads
	dummyLump := FileLump{}
	directorySize := int32(binary.Size(dummyLump)) * numwads
	directoryBuffer := bytes.NewBuffer(result.data[result.Header.InfoTableOffset : result.Header.InfoTableOffset+directorySize])

	result.Directory = make([]*FileLump, numwads)
	for i, _ := range result.Directory {
		var lumpInfo FileLump
		err = binary.Read(directoryBuffer, binary.LittleEndian, &lumpInfo)
		if err != nil {
			return
		}
		result.Directory[i] = &lumpInfo
	}
	return
}

func (w *Wad) ReadLump(lump *FileLump) ([]byte, error) {
	return w.data[lump.FilePos : lump.FilePos+lump.Size], nil
}
