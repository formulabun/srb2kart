package conversion

import "testing"

func TestMapIdToNumber(t *testing.T) {
  data := []struct{
    input string
    expected uint
    err error
  }{
    {"1", 1, nil},
    {"01", 1, nil},
    {"30", 30, nil},
    {"E4", 248, nil},
    {"GH", 333, nil},
    {"ZZ", 1035, nil},
    {"e4", 248, nil},
    {"gh", 333, nil},
    {"zz", 1035, nil},
    {"1A", 0, MapIdIncorrectFormat},
    {"123", 0, MapIdIncorrectFormat},
  }
  for _, d := range data {
    ret, err := MapIdToNumber(d.input)
    if err != d.err || ret != d.expected {
      t.Errorf("MapIdToNumber(%s) = (%d, %v) but got (%d, %v)",
      d.input, d.expected, d.err, ret, err)
    }
  }
}

func TestNumberToMapId(t *testing.T) {
	data := []struct {
		input    uint
		expected string
		err      error
	}{
		{1, "01", nil},
		{10, "10", nil},
		{30, "30", nil},
		{248, "E4", nil},
		{333, "GH", nil},
    {1035, "ZZ", nil},
    {1337, "", MapNumberOverflowError},
	}

	for _, d := range data {
		ret, err := NumberToMapId(d.input)
		if err != d.err || ret != d.expected {
			t.Errorf("NumberToMapId(%d) = (%s, %v) but got (%s, %v)",
				d.input, d.expected, d.err, ret, err)
		}
	}
}
