package main

var playPalTemplate = `package assets

import "go.formulabun.club/srb2kart/lump/palette"

// Generated file! Do not edit, but do commit

var PlayPal = palette.Palette{
{{- range .}}  { {{- .R}}, {{.G}}, {{.B}}, 0},
{{ end}}}
`


