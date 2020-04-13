package xrequest

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodGet:
			io.WriteString(writer, "Hello,GET")
		case http.MethodPost:
			io.Copy(writer, request.Body)
		case http.MethodPut:
		case http.MethodDelete:
		case http.MethodOptions:
		case http.MethodHead:
		case http.MethodTrace:
		case http.MethodPatch:
		case http.MethodConnect:

		}
	}))

	client := NewHttpClient(srv.Client(), "")

	t.Run("GET", func(t *testing.T) {
		response, err := client.Get(context.Background(), srv.URL)
		require.Nil(t, err, "err must be nil")
		body, err := response.String()
		require.Nil(t, err, "err must be nil")
		require.Equal(t, "Hello,GET", body)
	})

	t.Run("POST", func(t *testing.T) {
		reqBody, err := NewJsonBody(map[string]interface{}{"hello": "world", "age": 12})
		require.Nil(t, err, "err muse be nil")
		response, err := client.Post(context.Background(), srv.URL, reqBody)
		require.Nil(t, err, "err must be nil")
		body, err := response.String()
		require.Nil(t, err, "err must be nil")
		require.Equal(t, `{"age":12,"hello":"world"}`, body)
	})

	t.Run("PUT", func(t *testing.T) {

	})

	t.Run("DELETE", func(t *testing.T) {

	})

	t.Run("OPTIONS", func(t *testing.T) {

	})

	t.Run("HEAD", func(t *testing.T) {

	})

	t.Run("TRACE", func(t *testing.T) {

	})

	t.Run("PATCH", func(t *testing.T) {

	})

	t.Run("CONNECT", func(t *testing.T) {

	})
}

func TestInterceptor(t *testing.T) {
}

func TestBody(t *testing.T) {
	t.Run("json", func(t *testing.T) {

	})
	t.Run("xml", func(t *testing.T) {
		b, err := NewXmlBody(map[string]interface{}{"page": 1, "token": "abcd1234"})
		require.Nilf(t, err, "err must be nil, but got %s", err)
		require.Equal(t, "text/xml", b.ContentType())
		buf := make([]byte, 100)
		n, err := b.Read(buf)
		require.Nil(t, err, "err must be nil")
		require.Equal(t, "page=1&token=abcd1234", string(buf[:n]))
	})
	t.Run("url-encoded", func(t *testing.T) {
		b, err := NewUrlEncodedBody(url.Values{"page": []string{"1"}, "token": []string{"abcd1234"}})
		require.Nil(t, err, "err must be nil")
		require.Equal(t, "application/x-www-form-urlencoded", b.ContentType())
		buf := make([]byte, 100)
		n, err := b.Read(buf)
		require.Nil(t, err, "err must be nil")
		require.Equal(t, "page=1&token=abcd1234", string(buf[:n]))
	})
}
