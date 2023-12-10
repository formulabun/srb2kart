package network

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestScanFile(t *testing.T) {
	testData := []struct {
		data    []byte
		advance int
		token   []byte
		isErr   bool
	}{
		{
			[]byte{1, 0, 90, 224, 183, 0, 86, 83, 95, 83, 82, 66, 50, 84, 104, 105, 99, 99, 46, 112, 107, 51, 0, 90, 10, 80, 195, 31, 160, 113, 229, 102, 218, 100, 199, 125, 186, 237, 205},
			0,
			[]byte{},
			true,
		},
		{
			[]byte{17, 123, 14, 1, 0, 75, 76, 95, 72, 79, 83, 84, 77, 79, 68, 95, 86, 49, 54, 46, 112, 107, 51, 0, 230, 79, 249, 250, 150, 185, 117, 162, 192, 177, 220, 216, 113, 154, 102, 125, 17, 5, 71, 80, 0, 75, 66, 95, 77, 111, 117, 110, 116, 97, 105, 110, 83, 104, 111},
			5 + 19 + 16,
			[]byte{17, 123, 14, 1, 0, 75, 76, 95, 72, 79, 83, 84, 77, 79, 68, 95, 86, 49, 54, 46, 112, 107, 51, 0, 230, 79, 249, 250, 150, 185, 117, 162, 192, 177, 220, 216, 113, 154, 102, 125},
			false,
		},
	}

	for _, td := range testData {
		adv, tok, err := scanFile(td.data, false)
		if adv != td.advance || !slices.Equal(tok, td.token) || (td.isErr && err == nil) || (!td.isErr && err != nil) {
			t.Fatalf("scanFile(%v) =\n%v, %v, %v but expected\n%v, %v, %v == (err == nil)\n", td.data, adv, tok, err, td.advance, td.token, td.isErr)
		}
	}
}
