package assets

import "testing"

func TestPlayPalGeneration(T *testing.T) {
	if len(PlayPal) != 256 {
		T.Fatal("PLAYPAL platte isn't 256 colors big")
	}
}
