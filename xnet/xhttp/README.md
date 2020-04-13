# XHttp
增强了net/http包的能力，提供了中间件模式以及更简单的request工具。

[xrequest](./xrequest/README.md) http 请求包

## Middleware
```go
package xhttp

type Middleware interface {
	Next(h http.Handler) http.Handler
}

type MiddlewareFn func(h http.Handler) http.Handler

func (fn MiddlewareFn) Next(h http.Handler) http.Handler { return fn(h) }
```
`Middleware` 是服务端处理http请求的中间件，我们可以使用该接口来对http处理做一些额外的Hook，比如打日志，压缩response，路由中转等操作。

一个简单的日志中间件如下：
```go
func logHandler(w io.StringWriter) Middleware {
	return MiddlewareFn(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			now := time.Now()
			h.ServeHTTP(writer, request)
			w.WriteString(fmt.Sprintf("request method: %s, path: %s, latency: %s\n", request.Method, request.URL.Path, time.Since(now)))
		})
	})
}
```
我们在请求处理前记录下当前时间，在请求后记录下当前时间，然后做减法可以得到请求的耗时，最后进行记录。

## Headers
HTTP协议中定义了大量的Header Key，这里罗列了大量的通用Header Key用来简化应用开发。

## Method 与 Status
默认的Method和Status均为字符串，少了很多有用的特性，在这里，我们对他们进行了重定义，并且赋予了他们很多新的方法。

```go
package xhttp
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
```

### func (m Method) HasRequestBody() bool

是否可以有Request Body

### func (m Method) HasResponseBody() bool

是否可以用Response Body

### func (m Method) Safe() bool

是否安全

### func (m Method) IsIdempotent() bool

是否幂等

### func (m Method) Cacheable() bool

是否可以被缓存

### func (s *Status) Code() int
获取Status Code

### func (s *Status) Msg() string
获取Status Msg

### func (s *Status) IsOk() bool
是否是正确状态

### func (s *Status) IsClientError() bool
是否为客户端错误

### func (s *Status) IsServerError() bool
是否为服务端错误
