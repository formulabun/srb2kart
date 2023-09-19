package addons

import (
	"testing"

	"go.formulabun.club/srb2kart/pk3"
	"go.formulabun.club/srb2kart/wad"
)

func TestWadImplementsAddon(t *testing.T) {
	var w interface{} = &wad.Wad{}
	_, ok := w.(Addon)
	if !ok {
		t.Fatal("wad does not implement addon interface")
	}
}

func TestPk3ImplementsAddon(t *testing.T) {
	var p interface{} = &pk3.Pk3{}
	_, ok := p.(Addon)
	if !ok {
		t.Fatal("wad does not implement addon interface")
	}
}
