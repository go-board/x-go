# XRequest

## Request
`http.Request`的包装。
```go
package xrequest

type Request struct {
	R *http.Request
}
```
### RequestOption
简化对http.Request的修改。
```go
package xrequest
type RequestOption func(req *http.Request)
```
分别提供了以下的实现：
- AddHeader(key string, value string) RequestOption
- AddHeaderMap(h http.Header) RequestOption
- WithHeader(key string, value string) RequestOption
- WithHeaderMap(h http.Header) RequestOption
- WithForm(f url.Values) RequestOption
- WithContentType(contentType string) RequestOption
- AddCookie(cookie *http.Cookie) RequestOption
- WithRequestBody(body RequestBody) RequestOption

### RequestBody
请求体的抽象
```go
package xrequest

type RequestBody interface {
	io.Reader
	ContentType() string
	ContentEncoding() *string
}
```
默认提供了如下实现：
- JsonBody 序列化Json数据
- XmlBody 序列化Xml数据
- UrlEncodedBody 序列化UrlEncoded数据
- BinaryBody 序列化二进制数据
- GzipBody 压缩RequestBody

## Response
`http.Response`的包装
```go
type Response struct {
	Response *http.Response
	status   *xhttp.Status
}
```
## Client
`http.Client`的包装.
```go
package xrequest

import "net/http"

// RoundTripperFn implement http.RoundTripper for convenient usage.
type RoundTripperFn func(request *http.Request) (*http.Response, error)

func (fn RoundTripperFn) RoundTrip(request *http.Request) (*http.Response, error) { return fn(request) }

// Interceptor is interceptor that can do more work before/after an request
type Interceptor interface {
	Next(fn http.RoundTripper) http.RoundTripper
}

// InterceptorFn implement Interceptor for convenient usage.
type InterceptorFn func(rt http.RoundTripper) http.RoundTripper

func (fn InterceptorFn) Next(rt http.RoundTripper) http.RoundTripper { return fn(rt) }

type Client struct {
	client       *http.Client
	interceptors []Interceptor
	baseHost     string
}
```

### Methods
- func Head(ctx context.Context, url string, options ...RequestOption) (*Response, error)
- func Connect(ctx context.Context, url string, options ...RequestOption) (*Response, error)
- func Options(ctx context.Context, url string, options ...RequestOption) (*Response, error)
- func Trace(ctx context.Context, url string, options ...RequestOption) (*Response, error)
- func Get(ctx context.Context, url string, options ...RequestOption) (*Response, error)
- func Post(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error)
- func Put(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error)
- func Patch(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error)
- func Delete(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error)

### Interceptor
拦截器，在http请求前/后做一些额外的工作。比如日志，重试，负载均衡等操作。
简单的日志重试:
```go
package xrequest

func RetryOnStatusCode(codes ...int) InterceptorFn {
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(request *http.Request) (response *http.Response, err error) {
			for i := 0; i < 3; i++ {
				response, err = rt.RoundTrip(request)
				if err != nil || (response != nil && xslice.ContainsInt(codes, response.StatusCode)) {
					continue
				}
				return
			}
			return
		})
	}
}
```
