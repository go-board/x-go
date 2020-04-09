package xrequest

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-board/x-go/xnet/xhttp"
)

// Response is wrapper of http.Response with extra convenient methods.
type Response struct {
	Response *http.Response
	status   *xhttp.Status
}

func newResponse(r *http.Response) *Response {
	return &Response{Response: r, status: xhttp.NewStatus(r.StatusCode, r.Status)}
}

// Header return immutable response header which clone the original header.
func (r *Response) Header() http.Header { return r.Response.Header.Clone() }

// HeaderMut return mutable Response header.
func (r *Response) HeaderMut() http.Header { return r.Response.Header }

// ContentType return response content-type in header.
func (r *Response) ContentType() string { return r.Response.Header.Get("Content-Type") }

// ContentLength return response content-length in header.
func (r *Response) ContentLength() int64 { return r.Response.ContentLength }

// Status return response status.
func (r *Response) Status() *xhttp.Status { return r.status }

// Body return response body interface.
func (r *Response) Body() io.Reader { return r.Response.Body }

// Close close response body.
func (r *Response) Close() error { return r.Response.Body.Close() }

// JSON unmarshal response body data into v.
func (r *Response) JSON(v interface{}) error {
	d := json.NewDecoder(r.Response.Body)
	d.UseNumber()
	return d.Decode(v)
}

// XML unmarshal response body data into v.
func (r *Response) XML(v interface{}) error {
	return xml.NewDecoder(r.Response.Body).Decode(v)
}

// Bytes return response body in []byte format.
func (r *Response) Bytes() ([]byte, error) {
	defer r.Close()
	return ioutil.ReadAll(r.Response.Body)
}

// String return response body in string format.
func (r *Response) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Cookies retrieve response all cookies.
func (r *Response) Cookies() []*http.Cookie {
	return r.Response.Cookies()
}

// Cookie retrieve response cookie with name.
func (r *Response) Cookie(name string) (*http.Cookie, bool) {
	cookies := r.Response.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie, true
		}
	}
	return nil, false
}

// DownloadFile write response body to file.
func (r *Response) DownloadFile(name string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	defer r.Close()
	_, err = io.Copy(f, r.Response.Body)
	return err
}
