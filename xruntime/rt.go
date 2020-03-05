package xruntime

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
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
var gogcOnce sync.Once

// SetGoGc sets the garbage collection target percentage
// capture initial gc percentage first
// if gc > 0, set gc percentage to gc.
// else restore initial gc percentage and return.
func SetGoGc(gc int) {
	gogcOnce.Do(func() {
		gogc = debug.SetGCPercent(0)
	})
	if gc > 0 {
		debug.SetGCPercent(gc)
	} else {
		debug.SetGCPercent(gogc)
	}
}
