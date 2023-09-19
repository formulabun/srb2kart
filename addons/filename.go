package addons

import (
	"strconv"
	"strings"
	"unicode"
)

type AddonType uint8

const (
	KartFlag AddonType = 1 << iota
	SinglePlayerFlag
	MatchFlag
	RaceFlag
	FlagFlag
	BattleFlag
	CharFlag
	LuaFlag
)

func GetAddonType(filename string) AddonType {
	if filename == "bonuschars.kart" {
		return KartFlag | CharFlag
	}
	var result AddonType = 0
	for _, c := range filename {
		switch unicode.ToLower(c) {
		case 'k', 'x':
			result |= KartFlag
		case 's':
			result |= SinglePlayerFlag
		case 'm':
			result |= MatchFlag
		case 'r':
			result |= RaceFlag
		case 'f':
			result |= FlagFlag
		case 'b':
			result |= BattleFlag
		case 'c':
			result |= CharFlag
		case 'l':
			result |= LuaFlag
		case '_':
			return result
		}
	}
	return result
}

// Get the addon version from the filename as an array of numbers.
// This function does not catch non conforming version formats
func GetAddonVersion(filename string) []uint {
	lastSep := strings.LastIndexAny(filename, "-_")
	result := make([]uint, 0, 0)
	if lastSep < 0 {
		return result
	}
	i := lastSep + 1
	var digit uint
	for i < len(filename) {
		l := unicode.ToLower(rune(filename[i]))
		switch {
		case unicode.IsDigit(l):
			digit *= 10
			v, _ := strconv.Atoi(string(l))
			digit += uint(v)
			i++
		default:
			if unicode.IsDigit(rune(filename[i-1])) {
				result = append(result, digit)
			}
			digit = 0
			i++
		}
	}
	return result
}
