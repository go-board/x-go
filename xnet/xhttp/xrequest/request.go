package xrequest

import (
	"context"
	"net/http"

	"github.com/go-board/x-go/xnet/xhttp"
)

// Client is a client to perform http request and retrieve http response.
// Field client is *http.Client that real perform request.
// Field interceptors is chain of Interceptor to hook client.
type Client struct {
	client       *http.Client
	interceptors []Interceptor
}

var DefaultClient = &Client{
	client: http.DefaultClient,
}

// NewHttpClient create new client.
func NewHttpClient(client *http.Client, interceptors ...Interceptor) *Client {
	// try set client.Transport first
	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}
	// wrapper client.Transport
	for _, interceptor := range interceptors {
		client.Transport = interceptor.Next(client.Transport)
	}
	return &Client{client: client, interceptors: interceptors}
}

func (c *Client) newRequest(ctx context.Context, method string, url string, body xhttp.RequestBody, options ...RequestOption) (*http.Request, error) {
	r, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(r)
	}
	return r, nil
}

func (c *Client) doRequest(req *http.Request) (*Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return newResponse(response), nil
}

// Head do HEAD request.
func (c *Client) Head(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Connect do CONNECT request.
func (c *Client) Connect(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodConnect, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Options do OPTIONS request.
func (c *Client) Options(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodConnect, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Trace do TRACE request.
func (c *Client) Trace(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodTrace, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Get do GET request.
func (c *Client) Get(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodGet, url, nil, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Post do POST request.
func (c *Client) Post(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodPost, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return c.doRequest(r)
}

// Put do PUT request.
func (c *Client) Put(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodPut, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return c.doRequest(r)
}

// Patch do PATCH request.
func (c *Client) Patch(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodPatch, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return c.doRequest(r)
}

// Delete do DELETE request.
func (c *Client) Delete(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, http.MethodDelete, url, body, options...)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", body.ContentType())
	return c.doRequest(r)
}

// Head do HEAD request.
func Head(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return DefaultClient.Head(ctx, url, options...)
}

// Connect do CONNECT request.
func Connect(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return DefaultClient.Connect(ctx, url, options...)
}

// Options do OPTIONS request.
func Options(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return DefaultClient.Options(ctx, url, options...)
}

// Trace do TRACE request.
func Trace(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return DefaultClient.Trace(ctx, url, options...)
}

// Get do GET request.
func Get(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return DefaultClient.Get(ctx, url, options...)
}

// Post do POST request.
func Post(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Post(ctx, url, body, options...)
}

// Put do PUT request.
func Put(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Put(ctx, url, body, options...)
}

// Patch do PATCH request.
func Patch(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Patch(ctx, url, body, options...)
}

// Delete do DELETE request.
func Delete(ctx context.Context, url string, body xhttp.RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Delete(ctx, url, body, options...)
}
