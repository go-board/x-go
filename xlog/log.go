package xlog

import (
	"fmt"
	"io"
)

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

func Writer(w io.Writer) Logger {
	return &writerLogger{w: w}
}

type nopLogger struct{}

func (nopLogger) Print(args ...interface{})                 {}
func (nopLogger) Printf(format string, args ...interface{}) {}

func Nop() Logger { return nopLogger{} }
