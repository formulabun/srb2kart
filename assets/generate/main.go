package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"go.formulabun.club/functional/strings"
	"go.formulabun.club/srb2kart/lump/palette"
	"go.formulabun.club/srb2kart/wad"
)

var path = "/usr/share/games/SRB2Kart-v1.6/srb2.srb"

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func generatePlayPalPalette(lump []byte) {
	var rawData [256]struct {
		R, G, B byte
	}

	buff := bytes.NewBuffer(lump)
	palette, err := palette.Convert(buff)
	handleErr(err)

	templ, err := template.New("PLAYPAL").Parse(playPalTemplate)
	handleErr(err)
	outFile, err := os.OpenFile("playpal.go", os.O_RDWR|os.O_CREATE, 0644)
	handleErr(err)
	err = templ.Execute(outFile, palette)
	handleErr(err)
}

func main() {
	file, err := os.Open(path)
	handleErr(err)
	wadFile, err := wad.ReadWad(file)
	handleErr(err)

	var res *wad.FileLump

	for _, w := range wadFile.Directory {
		if strings.SafeNullTerminated(w.Name[:]) == "PLAYPAL" {
			res = w
			break
		}
	}
	lump, err := wad.ReadLump(file, res)
	generatePlayPalPalette(lump)
}
