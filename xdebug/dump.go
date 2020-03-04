package xdebug

import (
	"io"

	"github.com/davecgh/go-spew/spew"
)

func Debug(a ...interface{}) {
	spew.Dump(a...)
}

func FDebug(w io.Writer, a ...interface{}) {
	spew.Fdump(w, a...)
}

func SDebug(a ...interface{}) string {
	return spew.Sdump(a...)
}
