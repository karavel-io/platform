package logger

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
)

const (
	debugPrefix = "[DEBUG]"
	infoPrefix  = "[INFO]"
	warnPrefix  = "[DEBUG]"
	errorPrefix = "[ERROR]"
)

type Logger interface {
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Warn(a ...interface{})
	Warnf(format string, a ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Fatal(a ...interface{})
	Fatalf(format string, a ...interface{})
	Writer() io.Writer
	Level() Level
	SetLevel(lvl Level)
	SetPalette(palette Palette)
	SetColors(active bool)
}

type logger struct {
	w       io.Writer
	palette palette
	lvl     Level
	colors  bool
}

func New(lvl Level) Logger {
	return &logger{
		w:       color.Error,
		palette: palettes[PaletteDefault],
		lvl:     lvl,
	}
}

func (l *logger) Writer() io.Writer {
	return l.w
}

func (l *logger) Level() Level {
	return l.lvl
}

func (l *logger) SetLevel(lvl Level) {
	l.lvl = lvl
}

func (l *logger) SetPalette(palette Palette) {
	l.palette = palettes[palette]
}

func (l *logger) SetColors(active bool) {
	if active {
		l.palette.debug.EnableColor()
		l.palette.info.EnableColor()
		l.palette.warn.EnableColor()
		l.palette.error.EnableColor()
	} else {
		l.palette.debug.DisableColor()
		l.palette.info.DisableColor()
		l.palette.warn.DisableColor()
		l.palette.error.DisableColor()
	}
	l.colors = active
}

func (l *logger) Debug(a ...interface{}) {
	l.output(LvlDebug, a...)
}

func (l *logger) Debugf(format string, a ...interface{}) {
	l.outputf(LvlDebug, format, a...)
}

func (l *logger) Info(a ...interface{}) {
	l.output(LvlInfo, a...)
}

func (l *logger) Infof(format string, a ...interface{}) {
	l.outputf(LvlInfo, format, a...)
}

func (l *logger) Warn(a ...interface{}) {
	l.output(LvlWarn, a...)
}

func (l *logger) Warnf(format string, a ...interface{}) {
	l.outputf(LvlWarn, format, a...)
}

func (l *logger) Error(a ...interface{}) {
	l.output(LvlError, a...)
}

func (l *logger) Errorf(format string, a ...interface{}) {
	l.outputf(LvlError, format, a...)
}

func (l *logger) Fatal(a ...interface{}) {
	l.Error(a...)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, a ...interface{}) {
	l.Errorf(format, a...)
	os.Exit(1)
}

func (l *logger) output(lvl Level, a ...interface{}) {
	if !IsLevelActive(l.lvl, lvl) {
		return
	}

	prefix := false
	if len(a) > 0 {
		prefix = true
	}

	var p func(a ...interface{}) string
	switch lvl {
	case LvlDebug:
		a = append([]interface{}{debugPrefix, " "}, a...)
		p = l.palette.debug.SprintFunc()
	case LvlInfo:
		if prefix && !l.colors {
			a = append([]interface{}{infoPrefix, " "}, a...)
		}
		p = l.palette.info.SprintFunc()
	case LvlWarn:
		a = append([]interface{}{warnPrefix, " "}, a...)
		p = l.palette.warn.SprintFunc()
	case LvlError:
		a = append([]interface{}{errorPrefix, " "}, a...)
		p = l.palette.error.SprintFunc()
	default:
		return
	}
	_, _ = fmt.Fprintln(l.w, p(a...))
}

func (l *logger) outputf(lvl Level, s string, a ...interface{}) {
	if !IsLevelActive(l.lvl, lvl) {
		return
	}

	prefix := false
	if s != "" {
		prefix = true
	}

	var p func(format string, a ...interface{}) string
	switch lvl {
	case LvlDebug:
		s = debugPrefix + " " + s
		p = l.palette.debug.SprintfFunc()
	case LvlInfo:
		if prefix && !l.colors {
			s = infoPrefix + " " + s
		}
		p = l.palette.info.SprintfFunc()
	case LvlWarn:
		s = warnPrefix + " " + s
		p = l.palette.warn.SprintfFunc()
	case LvlError:
		s = errorPrefix + " " + s
		p = l.palette.error.SprintfFunc()
	default:
		return
	}

	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += "\n"
	}

	_, _ = fmt.Fprint(l.w, p(s, a...))
}
