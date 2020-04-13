package xdebug

import (
	"net/http"
	"runtime"
)

// gc is a http handler which will run runtime.GC when calling
type gc struct{}

func (gc) ServeHTTP(http.ResponseWriter, *http.Request) { runtime.GC() }

// GcHandler return gc handler
func GcHandler() http.Handler { return gc{} }
