package logger

type Level int

const (
	LvlDebug Level = iota
	LvlInfo
	LvlWarn
	LvlError
)

func IsLevelActive(current Level, wanted Level) bool {
	return current <= wanted
}
