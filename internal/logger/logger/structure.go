package logger_impl

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kyogai2281337/cns_eljur/internal/logger"
)

type Logger struct {
	level          logger.LogLevel
	Destinations   []logger.LogDest
	defaultTimeout time.Duration
}

func NewLogger(level logger.LogLevel, timeout int) *Logger {
	dur := time.Millisecond * time.Duration(timeout)
	return &Logger{
		level:          level,
		Destinations:   make([]logger.LogDest, 0),
		defaultTimeout: dur,
	}
}

func (l *Logger) AddDest(dest ...logger.LogDest) {
	l.Destinations = append(l.Destinations, dest...)
}

func (l *Logger) write(level logger.LogLevel, data string) error {
	errChan := make(chan error)
	ctx, cancel := context.WithTimeout(context.Background(), l.defaultTimeout)
	defer cancel()
	wg := new(sync.WaitGroup)
	wg.Add(len(l.Destinations))

	for idx, ld := range l.Destinations {
		go func(ld logger.LogDest) {
			defer wg.Done()
			if err := ld.Write(level, data); err != nil {
				errChan <- fmt.Errorf("logging error in index %d: %v", idx, err)
			}
		}(ld)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var allErrs []error

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err, ok := <-errChan:
			if !ok {
				if len(allErrs) > 0 {
					return fmt.Errorf("multiple errors occured: %+v", allErrs)
				}
				return nil
			}
			if err != nil {
				allErrs = append(allErrs, err)
			}
		}
	}
}

func (l *Logger) Trace(data string) error {
	return l.write(logger.LTrace, data)
}

func (l *Logger) Info(data string) error {
	return l.write(logger.LInfo, data)
}

func (l *Logger) Warn(data string) error {
	return l.write(logger.LWarn, data)
}

func (l *Logger) Error(data string) error {
	return l.write(logger.LError, data)
}

func (l *Logger) Fatal(data string) error {
	return l.write(logger.LFatal, data)
}
