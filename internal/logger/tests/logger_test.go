package logger_test

import (
	"testing"

	"github.com/kyogai2281337/cns_eljur/internal/logger"
	logdest "github.com/kyogai2281337/cns_eljur/internal/logger/destinations"
	logger_impl "github.com/kyogai2281337/cns_eljur/internal/logger/logger"
)

func TestStdout(t *testing.T) {
	l := logger_impl.NewLogger(logger.LInfo, 200)
	l.AddDest(logdest.NewStdoutLogDest(logger.LInfo))
	err := l.Info("AmogusTesting")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFile(t *testing.T) {
	l := logger_impl.NewLogger(logger.LInfo, 250)
	l.AddDest(logdest.NewFileLogDest(logger.LInfo, "test.log"))
	err := l.Info("AmogusTesting_In_File!!!")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

// !    16016             85217 ns/op             986 B/op         20 allocs/op
// * On a weak one, maybe it`s OK
func BenchmarkBoth(b *testing.B) {
	l := logger_impl.NewLogger(logger.LWarn, 150)
	l.AddDest(logdest.NewFileLogDest(logger.LWarn, "test.log"))
	l.AddDest(logdest.NewStdoutLogDest(logger.LWarn))
	for i := 0; i < b.N; i++ {
		if err := l.Error("ErrorData"); err != nil {
			b.Logf("Error occured: %s", err.Error())
		}
	}

}
