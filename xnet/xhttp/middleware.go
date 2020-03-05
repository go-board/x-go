package xhttp

import (
	"net/http"
)

// Middleware is HTTP processor middleware, called before next call.
// Middleware usually used by call some middlewares one by one.
// Example use:
//
// A --> before
//    B --> before
// 	     C --> before
//	        user http handler
//	     C --> after
//	  B --> after
// A --> after
type Middleware interface {
	Next(h http.Handler) http.Handler
}

// MiddlewareFn support wrap function with same signature as Middleware.
type MiddlewareFn func(h http.Handler) http.Handler

func (fn MiddlewareFn) Next(h http.Handler) http.Handler {
	return fn(h)
}
