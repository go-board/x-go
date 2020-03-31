package xrequest

import (
	"context"
	"net/http"
)

func newRequest(ctx context.Context, method string, url string, body Body, options ...RequestOption) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(r)
	}
	return r, nil
}

func doRequest(client *http.Client, req *http.Request) (*Response, error) {
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return newResponse(response), nil
}

// Head do HEAD request.
func Head(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return doRequest(http.DefaultClient, r)
}

// Connect do CONNECT request.
func Connect(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodConnect, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return doRequest(http.DefaultClient, r)
}

// Options do OPTIONS request.
func Options(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodConnect, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return doRequest(http.DefaultClient, r)
}

// Trace do TRACE request.
func Trace(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodTrace, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return doRequest(http.DefaultClient, r)
}

// Get do GET request.
func Get(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return doRequest(http.DefaultClient, r)
}

// Post do POST request.
func Post(ctx context.Context, url string, body Body, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodPost, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return doRequest(http.DefaultClient, r)
}

// Put do PUT request.
func Put(ctx context.Context, url string, body Body, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodPut, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return doRequest(http.DefaultClient, r)
}

// Patch do PATCH request.
func Patch(ctx context.Context, url string, body Body, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodPatch, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return doRequest(http.DefaultClient, r)
}

// Delete do DELETE request.
func Delete(ctx context.Context, url string, body Body, options ...RequestOption) (*Response, error) {
	r, err := newRequest(ctx, http.MethodDelete, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return doRequest(http.DefaultClient, r)
}
