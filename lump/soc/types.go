package soc

import "strings"

type Soc []Block

type Block struct {
	Header     Header
	Properties map[string]string
}

type Header struct {
	Type, Name string
}

func (b Block) IsLevel() bool {
	return strings.ToLower(b.Header.Type) == "level"
}
