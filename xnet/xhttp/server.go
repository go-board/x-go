package xhttp

import (
	"net"
	"net/http"

	"github.com/go-board/x-go/xnet"
)

// ServeLimitedServer run on a limited listener which limit maxium concurrency connection.
func ServeLimitedServer(server *http.Server, address string, n int) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	limitedListener := xnet.LimitedListener(ln, n)
	return server.Serve(limitedListener)
}

// ServeTLSLimitedServer run on a limited tls listener which limit maxium concurrency connection.
func ServeTLSLimitedServer(server *http.Server, address string, n int, certFile string, keyFile string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	limitedListener := xnet.LimitedListener(ln, n)
	return server.ServeTLS(limitedListener, certFile, keyFile)
}
