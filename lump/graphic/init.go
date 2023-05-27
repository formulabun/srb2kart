package graphic

import (
	"image"
	"image/color"
	"io"

	"go.formulabun.club/srb2kart/assets"
)

func Decode(r io.Reader) (image.Image, error) {
	return DecodeWithPalette(r, assets.PlayPal)
}

func DecodeWithPalette(r io.Reader, p color.Palette) (image.Image, error) {
	var head header
	var result image.Image
	err := readHeader(r, &head)
	if err != nil {
		return result, err
	}

	err = readImage(r, &head)
	if err != nil {
		return result, err
	}

	result, err = graphicToImage(&head, p)

	return result, err
}

func decodeConfig(r io.Reader) (image.Config, error) {
	var head header
	var result image.Config
	err := readHeader(r, &head)
	if err != nil {
		return result, err
	}

	result = image.Config{
		assets.PlayPal,
		int(head.width),
		int(head.height),
	}
	return result, nil
}

func init() {
	image.RegisterFormat("lmp", "?", Decode, decodeConfig)
}
