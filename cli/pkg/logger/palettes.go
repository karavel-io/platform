package logger

import "github.com/fatih/color"

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
		debug: color.New(color.FgGreen),
		info:  color.New(),
		warn:  color.New(color.FgYellow),
		error: color.New(color.FgRed),
	},
}
