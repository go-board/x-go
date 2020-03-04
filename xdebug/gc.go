package xdebug

import (
	"net/http"
	"runtime"
)

type gc struct{}

func (gc) ServeHTTP(http.ResponseWriter, *http.Request) { runtime.GC() }
