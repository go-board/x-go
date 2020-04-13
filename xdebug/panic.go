package xdebug

import (
	"net/http"
)

type panicHandler struct{}

func (panicHandler) ServeHTTP(http.ResponseWriter, *http.Request) { panic("user panic") }

// PanicHandler return panic handler
func PanicHandler() http.Handler { return panicHandler{} }
