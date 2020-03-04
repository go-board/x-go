package xruntime

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// CallerName return current caller name
func CallerName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "???"
	}
	name := runtime.FuncForPC(pc).Name()
	items := strings.Split(name, "/")
	if len(items) > 0 {
		return items[len(items)-1]
	}
	return name
}

func Caller() string {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return "???"
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func BuildPath() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	return info.Path
}

// Stack get stacktrace with pretty format.
func Stack(pretty bool) []byte {
	stack := debug.Stack()
	// todo: clean no need data and format it.
	return stack
}

var gogc int

func SetGoGc(gc int) {
	gogc = debug.SetGCPercent(gc)
}

func RestoreGoGc() {
	debug.SetGCPercent(gogc)
}
