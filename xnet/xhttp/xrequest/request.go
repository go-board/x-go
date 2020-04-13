package xrequest

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-board/x-go/xnet/xhttp"
)

// Request wrap http.Request and give more method
type Request struct {
	R *http.Request
}

// NewRequest create new http request.
func NewRequest(ctx context.Context, url string, method string, body RequestBody, options ...RequestOption) (*Request, error) {
	req, err := http.NewRequestWithContext(ctx, url, method, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set(xhttp.HeaderContentType, body.ContentType())
	if contentEncoding := body.ContentEncoding(); contentEncoding != nil {
		req.Header.Set(xhttp.HeaderContentEncoding, *contentEncoding)
	}
	for _, option := range options {
		option(req)
	}
	return &Request{R: req}, nil
}

// WithJson set request body to json format
func (r *Request) WithJson(v interface{}) *Request {
	body, _ := NewJsonBody(v)
	return r.WithBody(body)
}

// WithXml set request body to xml format
func (r *Request) WithXml(v interface{}) *Request {
	body, _ := NewXmlBody(v)
	return r.WithBody(body)
}

// WithBody set request body
func (r *Request) WithBody(body RequestBody) *Request {
	r.R.Body = ioutil.NopCloser(body)
	return r
}

// Fetch do http request and retrieve response
func (r *Request) Fetch(ctx context.Context, client *Client) (*Response, error) {
	r.R = r.R.WithContext(ctx)
	return client.doRequest(r.R)
}

// BodyJson retrieve body into json data
func (r *Request) BodyJson(v interface{}) error {
	if !xhttp.Method(r.R.Method).HasRequestBody() {
		return errorBodyNotAllowed(r.R.Method)
	}
	if err := errorContentType(xhttp.MIMEApplicationJSON, r.R.Header); err != nil {
		return err
	}
	d := json.NewDecoder(r.R.Body)
	return d.Decode(v)
}

// BodyXml retrieve body into xml data
func (r *Request) BodyXml(v interface{}) error {
	if !xhttp.Method(r.R.Method).HasRequestBody() {
		return errorBodyNotAllowed(r.R.Method)
	}
	if err := errorContentType(xhttp.MIMETextXML, r.R.Header); err != nil {
		return err
	}
	return xml.NewDecoder(r.R.Body).Decode(v)
}

// BodyUrlEncoded retrieve body into url-encoded data
func (r *Request) BodyUrlEncoded() (url.Values, error) {
	err := r.R.ParseForm()
	if err != nil {
		return nil, err
	}
	if err := errorContentType(xhttp.MIMEApplicationForm, r.R.Header); err != nil {
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

func (j *JsonBody) ContentType() string { return xhttp.MIMEApplicationJSON }

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

func (x *XmlBody) ContentType() string { return xhttp.MIMETextXML }

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

func (u *UrlEncodedBody) ContentType() string { return xhttp.MIMEApplicationForm }

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
