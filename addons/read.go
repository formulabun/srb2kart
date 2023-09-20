package addons

import (
	"archive/zip"
	"bytes"
	"errors"
	"io"

	"go.formulabun.club/srb2kart/pk3"
	"go.formulabun.club/srb2kart/wad"
)

// Reads the file as a wad, if that fails, reads it as a pk3
func Read(file io.Reader) (Addon, error) {
	buf, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	buff := bytes.NewReader(buf)

	// try a wad
	wad, wadErr := wad.ReadWad(buff)
	if wadErr == nil {
		return &wad, nil
	}
	buff.Seek(0, io.SeekStart)

	// try a pk3
	zip, pk3Err := zip.NewReader(buff, int64(buff.Len()))
	if pk3Err == nil {
		return (*pk3.Pk3)(zip), nil
	}

	return nil, errors.Join(wadErr, pk3Err)
}
