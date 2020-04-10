package xhttp

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

// RequestBody is abstract of request body.
type RequestBody interface {
	io.Reader
	ContentType() string
	ContentEncoding() *string
}

// JsonBody used to encode data to json format
type JsonBody struct {
	Body interface{}
	r    io.Reader
}

// NewJsonBody make new RequestBody
func NewJsonBody(data interface{}) (*JsonBody, error) {
	b := &JsonBody{Body: data}
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	b.r = bytes.NewBuffer(buf)
	return b, nil
}

func (j *JsonBody) Read(p []byte) (n int, err error) { return j.r.Read(p) }

func (j *JsonBody) ContentType() string { return "application/json" }

func (j *JsonBody) ContentEncoding() *string { return nil }

// XmlBody used to encode data to xml format
type XmlBody struct {
	Body interface{}
	r    io.Reader
}

// NewXmlBody make new RequestBody
func NewXmlBody(data interface{}) (*XmlBody, error) {
	b := &XmlBody{Body: data}
	buf, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	b.r = bytes.NewBuffer(buf)
	return b, nil
}

func (x *XmlBody) Read(p []byte) (n int, err error) { return x.r.Read(p) }

func (x *XmlBody) ContentType() string { return "text/xml" }

func (x *XmlBody) ContentEncoding() *string { return nil }

// UrlEncodedBody used to encode data to url-encoded format
type UrlEncodedBody struct {
	Body url.Values
	r    io.Reader
}

// NewUrlEncodedBody make new RequestBody
func NewUrlEncodedBody(data url.Values) (*UrlEncodedBody, error) {
	b := &UrlEncodedBody{Body: data}
	encoded := data.Encode()
	b.r = strings.NewReader(encoded)
	return b, nil
}

func (u *UrlEncodedBody) Read(p []byte) (n int, err error) { return u.r.Read(p) }

func (u *UrlEncodedBody) ContentType() string { return "application/x-www-form-urlencoded" }

func (u *UrlEncodedBody) ContentEncoding() *string { return nil }

// BinaryBody used to encode data to binary format
type BinaryBody struct {
	Body        []byte
	r           io.Reader
	contentType string
}

// NewBinaryBody make new RequestBody
func NewBinaryBody(data []byte, contentType string) (*BinaryBody, error) {
	if contentType == "" {
		return nil, errors.New("err: nil content-type")
	}
	return &BinaryBody{Body: data, r: bytes.NewReader(data), contentType: contentType}, nil
}

func (b *BinaryBody) Read(p []byte) (n int, err error) { return b.r.Read(p) }

func (b *BinaryBody) ContentType() string { return b.contentType }

func (b *BinaryBody) ContentEncoding() *string { return nil }

var gzipEncoding = "gzip"

// GzipBody wrap low layer Body into gzip reader
type GzipBody struct {
	Body RequestBody
	r    io.Reader
}

// NewGzipBody make new RequestBody
func NewGzipBody(body RequestBody, level int) (*GzipBody, error) {
	buffer := &bytes.Buffer{}
	w, err := gzip.NewWriterLevel(buffer, level)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(w, body)
	if err != nil {
		return nil, err
	}
	return &GzipBody{
		Body: body,
		r:    buffer,
	}, nil
}

func (g *GzipBody) Read(p []byte) (n int, err error) { return g.r.Read(p) }

func (g *GzipBody) ContentType() string { return g.Body.ContentType() }

func (g *GzipBody) ContentEncoding() *string { return &gzipEncoding }
