package logdest

import (
	"bufio"
	"os"

	"github.com/kyogai2281337/cns_eljur/internal/logger"
)

type Stdout struct {
	WriteLevel logger.LogLevel
	buf        bufio.Writer
	wr         WriteRule
}

func NewStdoutLogDest(level logger.LogLevel) *Stdout {
	return &Stdout{
		WriteLevel: level,
		buf:        *bufio.NewWriter(os.Stdout),
		wr:         *NewWriteRule(),
	}
}

func (s *Stdout) SetLevel(level logger.LogLevel) {
	s.WriteLevel = level
}

func (s *Stdout) Write(level logger.LogLevel, message string) error {
	if s._levelCheck(level) {
		_, err := s.buf.WriteString(s.wr.Convert(level, message))
		if err != nil {
			return err
		}
	}
	if err := s.buf.Flush(); err != nil {
		return err
	}
	return nil
}

func (s *Stdout) _levelCheck(incoming logger.LogLevel) bool {
	return s.WriteLevel <= incoming
}
