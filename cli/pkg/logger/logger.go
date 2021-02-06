package logger

import (
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
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
}

type logger struct {
	w       io.Writer
	palette palette
	lvl     Level
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

func (l *logger) Debug(a ...interface{}) {
	l.output(LvlDebug, append([]interface{}{"[DEBUG]", " "}, a...)...)
}

func (l *logger) Debugf(format string, a ...interface{}) {
	l.outputf(LvlDebug, "[DEBUG] "+format, a...)
}

func (l *logger) Info(a ...interface{}) {
	l.output(LvlInfo, a...)
}

func (l *logger) Infof(format string, a ...interface{}) {
	l.outputf(LvlInfo, format, a...)
}

func (l *logger) Warn(a ...interface{}) {
	l.output(LvlWarn, append([]interface{}{"[WARNING]", " "}, a...)...)
}

func (l *logger) Warnf(format string, a ...interface{}) {
	l.outputf(LvlWarn, "[WARNING] "+format, a...)
}

func (l *logger) Error(a ...interface{}) {
	l.output(LvlError, append([]interface{}{"[ERROR]", " "}, a...)...)
}

func (l *logger) Errorf(format string, a ...interface{}) {
	l.outputf(LvlError, "[ERROR] "+format, a...)
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

	var p func(a ...interface{}) string
	switch lvl {
	case LvlDebug:
		p = l.palette.debug.SprintFunc()
	case LvlInfo:
		p = l.palette.info.SprintFunc()
	case LvlWarn:
		p = l.palette.warn.SprintFunc()
	case LvlError:
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

	var p func(format string, a ...interface{}) string
	switch lvl {
	case LvlDebug:
		p = l.palette.debug.SprintfFunc()
	case LvlInfo:
		p = l.palette.info.SprintfFunc()
	case LvlWarn:
		p = l.palette.warn.SprintfFunc()
	case LvlError:
		p = l.palette.error.SprintfFunc()
	default:
		return
	}

	if len(s) == 0 || s[len(s)-1] != '\n' {
		s += "\n"
	}

	_, _ = fmt.Fprint(l.w, p(s, a...))
}
