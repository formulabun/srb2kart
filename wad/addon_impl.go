package wad

import (
	"bytes"
	"io"
	"strings"

	"go.formulabun.club/srb2kart/lump/soc"
)

func (w *Wad) LumpNames() []string {
	res := make([]string, len(w.Directory))
	for i, f := range w.Directory {
		res[i] = string(f.Name[:])
	}
	return res
}

func (w *Wad) Lumps() ([]io.Reader, error) {
	res := make([]io.Reader, len(w.Directory))
	for i, l := range w.Directory {
		content, err := w.ReadLump(l)
		if err != nil {
			return res, err
		}
		res[i] = bytes.NewBuffer(content)
	}
	return res, nil
}

func (w *Wad) SocNames() []string {
	res := make([]string, 0)
	for _, l := range w.Directory {
		lumpName := string(l.Name[:])
		if isSocLump(l) {
			res = append(res, lumpName)
		}
	}
	return res
}

func (w *Wad) Soc(socName string) (soc.Soc, error) {
	var lump *FileLump
	for _, lump = range w.Directory {
		if string(lump.Name[:]) == socName {
			break
		}
	}

	data, err := w.ReadLump(lump)
	return soc.ParseSoc(data), err
}

func (w *Wad) Socs() ([]soc.Soc, error) {
	res := make([]soc.Soc, 0)
	for _, l := range w.Directory {
		if isSocLump(l) {
			d, err := w.ReadLump(l)
			if err != nil {
				return res, err
			}
			res = append(res, soc.ParseSoc(d))
		}
	}
	return res, nil
}

func isSocLump(lump *FileLump) bool {
	lumpName := strings.ToLower(string(lump.Name[:]))
	return lumpName == "maincfg\x00" || lumpName == "objctcfg" || lumpName[:4] == "soc_"
}
