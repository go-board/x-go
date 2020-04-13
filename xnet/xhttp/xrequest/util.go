package xrequest

import (
	"fmt"
	"net/http"

	"github.com/go-board/x-go/xnet/xhttp"
)

func errorBodyNotAllowed(method string) error {
	return fmt.Errorf("method %s don't take any body", method)
}

func errorContentType(expected string, header http.Header) error {
	contentType := header.Get(xhttp.HeaderContentType)
	if expected == contentType {
		return nil
	}
	return fmt.Errorf("expected content-type is %s, but header %s", expected, contentType)
}
