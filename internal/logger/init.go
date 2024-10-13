package logger

type LogLevel byte

const (
	LTrace LogLevel = iota
	LInfo
	LWarn
	LError
	LFatal
)

type LogDest interface {
	SetLevel(level LogLevel)
	Write(level LogLevel, data string) error // ? maybe i`ll add concatinations
}

type Logger interface {
	AddDest(dest LogDest)
	write(level LogLevel, data string) error
	// * Simplifiers:
	Trace(data string) error
	Info(data string) error
	Warn(data string) error
	Error(data string) error
	Fatal(data string) error
}
