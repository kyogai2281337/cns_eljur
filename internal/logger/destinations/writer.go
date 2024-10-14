package logdest

import (
	"strings"
	"time"

	"github.com/kyogai2281337/cns_eljur/internal/logger"
)

const (
	traceVal = "| TRACE | "
	infoVal  = "| INFO | "
	warnVal  = "| WARN | "
	errorVal = "| ERROR | "
	fatalVal = "| FATAL | "
	undefVal = "| undefined | "
)

type WriteRule struct {
	b strings.Builder
}

func NewWriteRule() *WriteRule {
	return &WriteRule{
		b: strings.Builder{},
	}
}

func (wr *WriteRule) Convert(level logger.LogLevel, data string) string {
	defer wr.b.Reset()

	switch level {
	case logger.LTrace:
		wr.b.WriteString(traceVal)
	case logger.LInfo:
		wr.b.WriteString(infoVal)
	case logger.LWarn:
		wr.b.WriteString(warnVal)
	case logger.LError:
		wr.b.WriteString(errorVal)
	case logger.LFatal:
		wr.b.WriteString(fatalVal)
	default:
		wr.b.WriteString(undefVal)
	}
	wr.b.WriteString(time.Now().Format("2006 Jan 2 15:04:05 "))
	wr.b.WriteString(data + "\n")
	return wr.b.String()
}
