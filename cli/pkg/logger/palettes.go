package logger

import (
	"github.com/fatih/color"
)

type palette struct {
	debug *color.Color
	info  *color.Color
	warn  *color.Color
	error *color.Color
}

type Palette string

const (
	PaletteDefault Palette = "default"
)

var palettes = map[Palette]palette{
	PaletteDefault: {
		debug: color.New(color.FgHiGreen),
		info:  color.New(),
		warn:  color.New(color.FgHiYellow),
		error: color.New(color.FgHiRed),
	},
}
