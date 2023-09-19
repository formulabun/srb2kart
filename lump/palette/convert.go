package palette

import (
	"encoding/binary"
	"image/color"
	"io"

	"go.formulabun.club/functional/array"
)

const PALETTE_SIZE = 256

type rawColor struct {
	R, G, B byte
}

type rawPalette [PALETTE_SIZE]rawColor

func Convert(r io.Reader) (color.Palette, error) {
	var rawData rawPalette

	err := binary.Read(r, binary.LittleEndian, &rawData)
	if err != nil {
		return nil, err
	}

	palette := array.Map[rawColor, color.Color](rawData[:], func(c rawColor) color.Color {
		return color.RGBA{c.R, c.G, c.B, 255}
	})
	return color.Palette(palette), nil
}
