package xrequest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-board/x-go/xnet/xhttp"
)

func errorBodyNotAllowed(method string) error {
	return fmt.Errorf("method %s don't take any body", method)
}

func errorContentType(expected string, header http.Header) error {
	contentType := header.Get(xhttp.HeaderContentType)
	if expected == contentType ||
		strings.Contains(contentType, expected) { // for application/json; charset=utf-8 case
		return nil
	}
	return fmt.Errorf("expected content-type is %s, but header %s", expected, contentType)
}
