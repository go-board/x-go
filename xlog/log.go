package xlog

import (
	"fmt"
	"io"
)

// Logger is abstraction of log behavior.
type Logger interface {
	Print(args ...interface{})
	Printf(format string, args ...interface{})
}

type writerLogger struct {
	w io.Writer
}

func (w *writerLogger) Print(args ...interface{}) {
	_, _ = w.w.Write([]byte(fmt.Sprint(args...)))
}

func (w *writerLogger) Printf(format string, args ...interface{}) {
	_, _ = w.w.Write([]byte(fmt.Sprintf(format, args...)))
}

// Writer return a Logger which wrap an io.Writer.
func Writer(w io.Writer) Logger {
	return &writerLogger{w: w}
}

// MultiWriter return a Logger which write data to multiply io.Writer.
func MultiWriter(ws ...io.Writer) Logger {
	if len(ws) == 0 {
		return Nop()
	} else if len(ws) == 1 {
		return Writer(ws[0])
	}
	return &writerLogger{w: io.MultiWriter(ws...)}
}

type nopLogger struct{}

func (nopLogger) Print(args ...interface{})                 {}
func (nopLogger) Printf(format string, args ...interface{}) {}

// Nop return a Logger which do nothing.
func Nop() Logger { return nopLogger{} }
