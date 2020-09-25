package xrequest

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// RequestOption modify *http.Request in a convenient way.
type RequestOption func(req *http.Request)

// AddHeader append key/value pair to header.
// Deprecated: use WithAddHeader replaced.
func AddHeader(key string, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Add(key, value)
	}
}

// AddHeaderMap append extra header data to existing header.
func AddHeaderMap(h http.Header) RequestOption {
	return func(req *http.Request) {
		for k, v := range h {
			for _, vv := range v {
				req.Header.Add(k, vv)
			}
		}
	}
}

// WithHeader replace or set existing header with key/value pair.
// Deprecated: use WithSetHeader replaced.
func WithHeader(key string, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// WithAddHeader append header value to key.
func WithAddHeader(key string, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Add(key, value)
	}
}

// WithSetHeader upsert header value to key.
func WithSetHeader(key string, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}

// WithHeaderMap replace or set existing header.
func WithHeaderMap(h http.Header) RequestOption {
	return func(req *http.Request) {
		req.Header = h
	}
}

// WithForm replace or set existing form.
func WithForm(f url.Values) RequestOption {
	return func(req *http.Request) {
		req.Form = f
	}
}

// WithContentType set `Content-Type` header.
func WithContentType(contentType string) RequestOption {
	return WithSetHeader("Content-Type", contentType)
}

// AddCookie append cookie to request.
func AddCookie(cookie *http.Cookie) RequestOption {
	return func(req *http.Request) {
		req.AddCookie(cookie)
	}
}

func WithRequestBody(body RequestBody) RequestOption {
	return func(req *http.Request) {
		req.Body = ioutil.NopCloser(body)
	}
}

// WithAddQuery add query param to url.
func WithAddQuery(key, value string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Add(key, value)
		req.URL.RawQuery = q.Encode()
	}
}

// WithSetQuery set query param to url.
func WithSetQuery(key, value string) RequestOption {
	return func(req *http.Request) {
		q := req.URL.Query()
		q.Set(key, value)
		req.URL.RawQuery = q.Encode()
	}
}
