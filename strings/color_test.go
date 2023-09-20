package strings

import "testing"

func TestRemoveColorCodes(t *testing.T) {
	testData := []struct {
		in, out string
	}{
		{"Ring Cup 1 - \x8CRev", "Ring Cup 1 - Rev"},
	}

	for _, td := range testData {
		calc := RemoveColorCodes(td.in)
		if calc != td.out {
			t.Errorf("RemoveColorCode(%#v) = %#v but expected %#v", td.in, calc, td.out)
		}
	}

}
