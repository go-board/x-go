package metadata

import (
	"net/http"
	"strings"
)

var xPrefixKey = "x-md-"

// SetHTTPPrefixKey update http header prefix with `key`, default is 'x-md-'
//
// 设置http header中metadata的前缀，默认是'x-md-'
func SetHTTPPrefixKey(key string) {
	xPrefixKey = key
}

func (md Metadata) HTTPHeader() http.Header {
	h := http.Header{}
	for k, v := range md {
		h.Add("x-md-"+k, v)
	}
	return h
}

// IncomingHTTPHeader retrieve Metadata from an HTTPHeader header
//
// 从HTTP头中取出Metadata
func IncomingHTTPHeader(h http.Header) Metadata {
	md := Metadata{}
	for k, vs := range h {
		if strings.HasPrefix(k, xPrefixKey) {
			md[k] = strings.Join(vs, ",")
		}
	}
	return md
}

// OutgoingHTTPHeader append Metadata to an HTTPHeader header
//
// 向HTTP头中加入Metadata
func OutgoingHTTPHeader(md Metadata, h http.Header) http.Header {
	for k, v := range md {
		h.Add(xPrefixKey+k, v)
	}
	return h
}

// IncomingHTTPRequest retrieve Metadata from an HTTPHeader request
//
// 从HTTP请求中取出Metadata
func IncomingHTTPRequest(r *http.Request) Metadata {
	md := Metadata{}
	for k, vs := range r.Header {
		if strings.HasPrefix(k, xPrefixKey) {
			md[k] = strings.Join(vs, ",")
		}
	}
	return md
}

// OutgoingHTTPRequest append Metadata to an HTTPHeader request
//
// 向HTTP请求中加入Metadata
func OutgoingHTTPRequest(md Metadata, r *http.Request) *http.Request {
	for k, v := range md {
		r.Header.Add(xPrefixKey+k, v)
	}
	return r
}

// IncomingHTTPResponse retrieve Metadata from an HTTPHeader response
//
// 从HTTP响应中取出Metadata
func IncomingHTTPResponse(r *http.Response) Metadata {
	md := Metadata{}
	for k, vs := range r.Header {
		if strings.HasPrefix(k, xPrefixKey) {
			md[k] = strings.Join(vs, ",")
		}
	}
	return md
}

// OutgoingHTTPResponse append Metadata to an HTTPHeader response
//
// 向HTTP响应中加入Metadata
func OutgoingHTTPResponse(md Metadata, r *http.Response) *http.Response {
	for k, v := range md {
		r.Header.Add(xPrefixKey+k, v)
	}
	return r
}
