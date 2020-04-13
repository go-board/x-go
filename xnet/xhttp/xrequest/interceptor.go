package xrequest

import (
	"compress/gzip"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/go-board/x-go/xnet/xhttp"
	"github.com/go-board/x-go/xslice"
)

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

// ComposeInterceptor compose interceptors to given http.RoundTripper
func ComposeInterceptor(rt http.RoundTripper, interceptors ...Interceptor) http.RoundTripper {
	if len(interceptors) == 0 {
		return rt
	}
	return ComposeInterceptor(interceptors[0].Next(rt), interceptors[1:]...)
}

// InjectHeader inject given header into request.
func InjectHeader(h http.Header) InterceptorFn {
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(req *http.Request) (*http.Response, error) {
			for k, v := range h {
				for _, vv := range v {
					req.Header.Add(k, vv)
				}
			}
			return rt.RoundTrip(req)
		})
	}
}

// Logging is Interceptor that log http request stats
func Logging(rt http.RoundTripper) http.RoundTripper {
	return RoundTripperFn(func(request *http.Request) (*http.Response, error) {
		before := time.Now()
		response, err := rt.RoundTrip(request)
		if err != nil {
			log.Printf("%s %s, latency: %s, status: %s\n", request.Method, request.URL.Path, time.Since(before), err)
		} else {
			log.Printf("%s %s, latency: %s, status: %s\n", request.Method, request.URL.Path, time.Since(before), response.Status)
		}
		return response, err
	})
}

// RetryOnStatusCode retry on return codes...
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

// RetryStrategy is strategy for http request
type RetryStrategy struct {
	Backoff     func(r *http.Request, i int) time.Duration
	MaxRetries  func(r *http.Request) int
	ShouldRetry func(r *http.Request, resp *http.Response, err error) bool
}

// RetryWithStrategy retry with given strategy.
func RetryWithStrategy(strategy RetryStrategy) InterceptorFn {
	if strategy.Backoff == nil {
		strategy.Backoff = func(r *http.Request, i int) time.Duration { return 0 }
	}
	if strategy.MaxRetries == nil {
		strategy.MaxRetries = func(r *http.Request) int {
			if retryStr := r.Header.Get("X-Max-Retries"); retryStr != "" {
				if retries, err := strconv.ParseInt(retryStr, 10, 64); err == nil {
					return int(retries)
				}
			}
			return 3
		}
	}
	if strategy.ShouldRetry == nil {
		strategy.ShouldRetry = func(r *http.Request, resp *http.Response, err error) bool {
			return err != nil || xslice.ContainsInt([]int{500}, resp.StatusCode)
		}
	}
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(request *http.Request) (response *http.Response, err error) {
			maxRetries := strategy.MaxRetries(request)
			for i := 0; i < maxRetries; i++ {
				response, err = rt.RoundTrip(request)
				if strategy.ShouldRetry(request, response, err) {
					continue
				}
				return
			}
			return
		})
	}
}

// RoundRobinProxy proxy request with round robin strategy to different server.
func RoundRobinProxy(hosts ...string) InterceptorFn {
	if len(hosts) == 0 {
		panic("empty hosts list")
	}
	var term uint64 = 0
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(request *http.Request) (*http.Response, error) {
			host := hosts[atomic.AddUint64(&term, 1)%uint64(len(hosts))]
			request.Host = host
			return rt.RoundTrip(request)
		})
	}
}

// RandomProxy proxy request with random strategy to different server.
func RandomProxy(hosts ...string) InterceptorFn {
	if len(hosts) == 0 {
		panic("empty hosts list")
	}
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(request *http.Request) (*http.Response, error) {
			request.Host = hosts[rand.Intn(len(hosts))]
			return rt.RoundTrip(request)
		})
	}
}

// GzipDecompressResponse decompress response body if possible.
func GzipDecompressResponse() InterceptorFn {
	return func(rt http.RoundTripper) http.RoundTripper {
		return RoundTripperFn(func(request *http.Request) (*http.Response, error) {
			response, err := rt.RoundTrip(request)
			if err != nil {
				return nil, err
			}
			if response.Header.Get(xhttp.HeaderContentEncoding) == "gzip" {
				r, err := gzip.NewReader(response.Body)
				if err != nil {
					return nil, err
				}
				response.Body = r
			}
			return response, err
		})
	}
}
