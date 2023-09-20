package addons

import (
	"io"

	"go.formulabun.club/srb2kart/lump/soc"
)

type Addon interface {
	Lumps() ([]io.Reader, error)
	LumpNames() []string

	// using the output from Socs or LumpNames
	Soc(socName string) (soc.Soc, error)
	Socs() ([]soc.Soc, error)
	SocNames() []string
}
