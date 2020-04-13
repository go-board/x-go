package xdebug

import (
	"io"

	"github.com/davecgh/go-spew/spew"
)

// Debug is pretty print version of fmt.Print
func Debug(a ...interface{}) {
	spew.Dump(a...)
}

// FDebug is pretty print version of fmt.Printf
func FDebug(w io.Writer, a ...interface{}) {
	spew.Fdump(w, a...)
}

// Debug is pretty print version of fmt.Sprint
func SDebug(a ...interface{}) string {
	return spew.Sdump(a...)
}
