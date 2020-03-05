package xhttp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func logHandler(w io.StringWriter) Middleware {
	return MiddlewareFn(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			now := time.Now()
			h.ServeHTTP(writer, request)
			w.WriteString(fmt.Sprintf("request method: %s, path: %s, latency: %s\n", request.Method, request.URL.Path, time.Since(now)))
		})
	})
}

func responseHandler(body string) Middleware {
	return MiddlewareFn(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			io.WriteString(writer, body)
		})
	})
}

func TestMiddleware(t *testing.T) {
	httpHandler := func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 100)
	}

	b := &bytes.Buffer{}
	logHandler(b).Next(http.HandlerFunc(httpHandler)).ServeHTTP(nil, &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/v2/api/user/info"}})
	require.Contains(t, b.String(), http.MethodGet)
	require.Contains(t, b.String(), "/v2/api/user/info")
}

func TestChainedMiddleware(t *testing.T) {
	httpHandler := func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 100)
	}
	b := &bytes.Buffer{}
	resp := httptest.NewRecorder()
	logIt := logHandler(b).Next(http.HandlerFunc(httpHandler))
	respIt := responseHandler("Hello,world").Next(logIt)
	respIt.ServeHTTP(resp, &http.Request{Method: http.MethodGet, URL: &url.URL{Path: "/v2/api/user/info"}})
	body, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err, "err must be nil")
	require.Equal(t, "Hello,world", string(body))
}
