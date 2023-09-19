package wad

import (
	"errors"
	"fmt"
)

func validateHeader(head *Header) error {
	switch string(head.Identification[:]) {
	case "IWAD":
	case "PWAD":
	case "ZWAD":
		break
	case "SDLL":
		return fmt.Errorf("Unsupported wad type: '%s'", head.Identification)
	default:
		return errors.New("Unknown wad type")
	}
	if head.NumWads < 0 || head.InfoTableOffset < 0 {
		return errors.New("Invalid value in wad header")
	}
	return nil
}
