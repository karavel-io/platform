// Copyright 2021 MIKAMAI s.r.l
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
