package xhttp

import (
	"fmt"
	"net/http"
)

// Method is http request method.
type Method string

const (
	MethodConnect Method = "CONNECT"
	MethodDelete         = "DELETE"
	MethodGet            = "GET"
	MethodHead           = "HEAD"
	MethodOptions        = "OPTIONS"
	MethodPatch          = "PATCH"
	MethodPost           = "POST"
	MethodPut            = "PUT"
	MethodTrace          = "TRACE"
)

// HasRequestBody indicate whether has a body in request or not.
func (m Method) HasRequestBody() bool {
	return m == MethodDelete ||
		m == MethodPatch ||
		m == MethodPost ||
		m == MethodPut
}

// HasResponseBody indicate whether has a body in response or not.
func (m Method) HasResponseBody() bool {
	return m == MethodConnect ||
		m == MethodDelete ||
		m == MethodGet ||
		m == MethodPost
}

// Safe indicate whether request is safe or not.
func (m Method) Safe() bool {
	return m == MethodGet ||
		m == MethodHead ||
		m == MethodOptions
}

// IsIdempotent indicate whether request is idempotent or not.
func (m Method) IsIdempotent() bool {
	return m == MethodDelete ||
		m == MethodGet ||
		m == MethodHead ||
		m == MethodOptions ||
		m == MethodPut ||
		m == MethodTrace
}

// Cacheable indicate whether client can cache response or not.
func (m Method) Cacheable() bool {
	return m == MethodGet ||
		m == MethodHead ||
		m == MethodPost
}

// Status is response status with a code and a msg string.
type Status struct {
	code int
	msg  string
}

// NewStatus create new status
func NewStatus(code int, msg string) *Status {
	return &Status{
		code: code,
		msg:  msg,
	}
}

// StatusOk return ok status
func StatusOk() *Status {
	return &Status{code: http.StatusOK, msg: http.StatusText(http.StatusOK)}
}

// StatusBadRequest is bad request -- 400
func StatusBadRequest() *Status {
	return &Status{code: http.StatusBadRequest, msg: http.StatusText(http.StatusBadRequest)}
}

// StatusServerInternalError is server error -- 500
func StatusServerInternalError() *Status {
	return &Status{code: http.StatusInternalServerError, msg: http.StatusText(http.StatusInternalServerError)}
}

// Code return status code.
func (s *Status) Code() int { return s.code }

// Msg return status msg.
func (s *Status) Msg() string { return s.msg }

// String return status string.
func (s *Status) String() string {
	return fmt.Sprintf("Response code %d, msg is %s", s.code, s.msg)
}

// IsOk check whether the request is successful.
func (s *Status) IsOk() bool { return s.code/100 == 2 }

// IsClientError check the error occur in client side.
func (s *Status) IsClientError() bool { return s.code/100 == 4 }

// IsServerError check the error occur in server side.
func (s *Status) IsServerError() bool { return s.code/100 == 5 }
