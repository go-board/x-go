package xhttp

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMethod(t *testing.T) {

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
