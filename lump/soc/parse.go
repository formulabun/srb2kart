package soc

import (
	"bufio"
	"bytes"
	"strings"
)

// sInit -> sBlock when reading a header line
// sBlock -> sInit when reading an empty line
const (
	// expecting to read a header
	sInit = iota
	// expecting to read a property line
	sBlock
)

const spaceCutSet = " "

func ParseSoc(data []byte) Soc {
	scanner := bufio.NewScanner(bytes.NewBuffer(data))
	scanner.Split(bufio.ScanLines)

	res := Soc{}
	state := sInit

	var i = -1
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), spaceCutSet)
		switch state {
		case sInit:
			if isEmptyLine(line) {
				continue
			} else {
				state = sBlock
				h := parseHeader(line)
				res = append(res, h)
				i += 1
			}
		case sBlock:
			if isEmptyLine(line) {
				state = sInit
			} else {
				k, v := parseProp(line)
				res[i].Properties[k] = v
			}
		}
	}
	return res
}

func isEmptyLine(line string) bool {
	return len(line) == 0 || line[0] == '#'
}

func parseHeader(line string) Block {
	before, _, _ := strings.Cut(line, "#")
	t, n, _ := strings.Cut(before, spaceCutSet)
	return Block{
		Header{
			strings.ToUpper(strings.Trim(t, spaceCutSet)),
			strings.Trim(n, spaceCutSet),
		},
		make(map[string]string),
	}
}

func parseProp(line string) (key, value string) {
	key, value, _ = strings.Cut(line, "=")
	return strings.ToUpper(strings.Trim(key, spaceCutSet)), strings.Trim(value, spaceCutSet)
}
