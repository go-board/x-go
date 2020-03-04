package xdebug

import (
	"net/http"
)

type panicHandler struct{}

func (panicHandler) ServeHTTP(http.ResponseWriter, *http.Request) { panic("user panic") }
