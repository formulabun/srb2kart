package conversion

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var MapNumberOverflowError = errors.New("Map number can't be bigger than 1035")
var MapIdIncorrectFormat = errors.New("Map id is incorrectly formatted")

// see http://wiki.srb2.org/wiki/Extended_map_number
func MapIdToNumber(id string) (uint, error) {
	l := len(id)
	switch {
	case l == 1:
		id = fmt.Sprintf("%s%s", "0", id)
	case l > 2:
		return 0, MapIdIncorrectFormat
	}
	var x, y byte
	var p, q int

	id = strings.ToUpper(id)
	x = id[0]
	y = id[1]
	if unicode.IsDigit(rune(x)) {
		if !unicode.IsDigit(rune(y)) {
			return 0, MapIdIncorrectFormat
		}
		n, err := strconv.Atoi(id)
		return uint(n), err
	}

	p = int(x - 'A')
	if unicode.IsDigit(rune(y)) {
		q = int(y - '0')
	} else {
		q = int(y-'A') + 10
	}
	return uint(100 + (36*p + q)), nil
}

func NumberToMapId(n uint) (string, error) {
	if n < 100 {
		return fmt.Sprintf("%02d", n), nil
	}
	if n > 1035 {
		return "", MapNumberOverflowError
	}
	var x, p, q int
	var a, b int
	x = int(n - 100)
	p = x / 36
	q = x - 36*p
	a = 'A' + p
	if q > 10 {
		b = 'A' + q - 10
	} else {
		b = '0' + q
	}

	return fmt.Sprintf("%c%c", a, b), nil
}
