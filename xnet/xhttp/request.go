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

func errorBodyNotAllowed(method string) error {
	return fmt.Errorf("method %s don't take any body", method)
}

func errorContentTypeMismatch(expected string, r *http.Request) error {
	contentType := r.Header.Get(HeaderContentType)
	if expected == contentType {
		return nil
	}
	return fmt.Errorf("expected content-type is %s, but got %s", expected, contentType)
}

// Request wrap http.Request and give more method
type Request struct {
	R *http.Request
}

func NewRequest(r *http.Request) (*Request, error) {
	contentEncoding := r.Header.Get(HeaderContentEncoding)
	// todo: support more compress option
	switch contentEncoding {
	case "gzip":
		reader, err := gzip.NewReader(r.Body)
		if err != nil {
			return nil, err
		}
		r.Body = reader
	default:
	}
	return &Request{R: r}, nil
}

func (r *Request) BodyJson(v interface{}) error {
	if !Method(r.R.Method).HasRequestBody() {
		return errorBodyNotAllowed(r.R.Method)
	}
	if err := errorContentTypeMismatch(MIMEApplicationJSON, r.R); err != nil {
		return err
	}
	d := json.NewDecoder(r.R.Body)
	return d.Decode(v)
}

func (r *Request) BodyXml(v interface{}) error {
	if !Method(r.R.Method).HasRequestBody() {
		return errorBodyNotAllowed(r.R.Method)
	}
	if err := errorContentTypeMismatch(MIMETextXML, r.R); err != nil {
		return err
	}
	return xml.NewDecoder(r.R.Body).Decode(v)
}

func (r *Request) BodyUrlEncoded() (url.Values, error) {
	err := r.R.ParseForm()
	if err != nil {
		return nil, err
	}
	if err := errorContentTypeMismatch(MIMEApplicationForm, r.R); err != nil {
		return nil, err
	}
	return r.R.PostForm, nil
}

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

func (j *JsonBody) ContentType() string { return MIMEApplicationJSON }

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

func (x *XmlBody) ContentType() string { return MIMETextXML }

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

func (u *UrlEncodedBody) ContentType() string { return MIMEApplicationForm }

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
		return nil, errors.New("err: empty content-type")
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
