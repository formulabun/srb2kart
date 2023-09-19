package main

var playPalTemplate = `package assets

import "image/color"

// Generated file! Do not edit, but do commit

var PlayPal = color.Palette{
{{- range .}}  color.RGBA{ {{- .R}}, {{.G}}, {{.B}}, {{.A}}},
{{ end}}}
`
