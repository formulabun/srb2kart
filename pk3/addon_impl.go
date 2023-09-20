package pk3

import (
	"errors"
	"io"
	"strings"

	"go.formulabun.club/srb2kart/lump/soc"
)

func (p *Pk3) Lumps() ([]io.Reader, error) {
	data := make([]io.Reader, len(p.File))
	for i, f := range p.File {
		var err error
		data[i], err = f.Open()
		if err != nil {
			return data, err
		}
	}
	return data, nil
}

func (p *Pk3) LumpNames() []string {
	data := make([]string, len(p.File))
	for i, f := range p.File {
		data[i] = f.Name
	}
	return data
}

func (p *Pk3) Soc(socName string) (soc.Soc, error) {
	for _, f := range p.File {
		if f.Name == socName {
			decompr, err := f.Open()
			defer decompr.Close()
			if err != nil {
				return soc.Soc{}, err
			}
			data, err := io.ReadAll(decompr)
			return soc.ParseSoc(data), err
		}
	}
	return soc.Soc{}, errors.New("File not Found")
}

func (p *Pk3) Socs() ([]soc.Soc, error) {
	data := make([]soc.Soc, 0)
	for _, f := range p.File {
		if isSocFile(f.Name) {
			decompr, err := f.Open()
			defer decompr.Close()
			if err != nil {
				return data, err
			}
			d, err := io.ReadAll(decompr)
			if err != nil {
				return data, err
			}
			data = append(data, soc.ParseSoc(d))
		}
	}
	return data, nil
}

func (p *Pk3) SocNames() []string {
	data := make([]string, 0)
	for _, f := range p.File {
		if isSocFile(f.Name) {
			data = append(data, f.Name)
		}
	}
	return data
}

func isSocFile(filename string) bool {
	sanitized := strings.ToLower(filename)
	return strings.HasPrefix(sanitized, "soc/") && !strings.HasSuffix(sanitized, "/")
}
