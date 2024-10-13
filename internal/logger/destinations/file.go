package logdest

import (
	"bufio"
	"os"

	"github.com/kyogai2281337/cns_eljur/internal/logger"
)

type File struct {
	WriteLevel logger.LogLevel
	buf        bufio.Writer
}

func NewFileLogDest(level logger.LogLevel, fileName string) *File {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	resp := &File{
		WriteLevel: level,
		buf:        *bufio.NewWriter(f),
	}

	return resp
}

func (f *File) SetLevel(level logger.LogLevel) {
	f.WriteLevel = level
}

func (f *File) Write(level logger.LogLevel, message string) error {
	if f._levelCheck(level) {
		_, err := f.buf.WriteString(message + "\n")
		if err != nil {
			return err
		}

		if err := f.buf.Flush(); err != nil {
			return err
		}
	}
	return nil
}

func (f *File) _levelCheck(incoming logger.LogLevel) bool {
	return f.WriteLevel <= incoming
}
