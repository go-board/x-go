package xrequest

import (
	"context"
	"net/http"
	"net/url"
)

// Client is a client to perform http request and retrieve http response.
// Field client is *http.Client that real perform request.
// Field interceptors is chain of Interceptor to hook client.
type Client struct {
	client       *http.Client
	interceptors []Interceptor
	baseHost     string
}

var DefaultClient = &Client{
	client: http.DefaultClient,
}

// NewHttpClient create new client.
func NewHttpClient(client *http.Client, baseHost string, interceptors ...Interceptor) *Client {
	// try set client.Transport first
	if client.Transport == nil {
		client.Transport = http.DefaultTransport
	}
	// wrapper client.Transport
	client.Transport = ComposeInterceptor(client.Transport, interceptors...)
	return &Client{client: client, interceptors: interceptors, baseHost: baseHost}
}

func (c *Client) newRequest(ctx context.Context, method string, rawurl string, body RequestBody, options ...RequestOption) (*http.Request, error) {
	if c.baseHost != "" {
		u, err := url.Parse(rawurl)
		if err == nil {
			rawurl = c.baseHost + u.RequestURI()
		}
	}
	r, err := http.NewRequestWithContext(ctx, method, rawurl, body)
	if err != nil {
		return nil, err
	}
	for _, option := range options {
		option(r)
	}
	if body != nil {
		r.Header.Set("Content-Type", body.ContentType())
		if encoding := body.ContentEncoding(); encoding != nil {
			r.Header.Set("Content-Encoding", *encoding)
		}
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

func (c *Client) perform(ctx context.Context, method string, rawurl string, body RequestBody, options ...RequestOption) (*Response, error) {
	r, err := c.newRequest(ctx, method, rawurl, body, options...)
	if err != nil {
		return nil, err
	}
	return c.doRequest(r)
}

// Head do HEAD request.
func (c *Client) Head(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodHead, url, nil, options...)
}

// Connect do CONNECT request.
func (c *Client) Connect(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodConnect, url, nil, options...)
}

// Options do OPTIONS request.
func (c *Client) Options(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodConnect, url, nil, options...)
}

// Trace do TRACE request.
func (c *Client) Trace(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodTrace, url, nil, options...)
}

// Get do GET request.
func (c *Client) Get(ctx context.Context, url string, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodGet, url, nil, options...)
}

// Post do POST request.
func (c *Client) Post(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodPost, url, body, options...)
}

// Put do PUT request.
func (c *Client) Put(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodPut, url, body, options...)
}

// Patch do PATCH request.
func (c *Client) Patch(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodPatch, url, body, options...)
}

// Delete do DELETE request.
func (c *Client) Delete(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return c.perform(ctx, http.MethodDelete, url, body, options...)
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
func Post(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Post(ctx, url, body, options...)
}

// Put do PUT request.
func Put(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Put(ctx, url, body, options...)
}

// Patch do PATCH request.
func Patch(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Patch(ctx, url, body, options...)
}

// Delete do DELETE request.
func Delete(ctx context.Context, url string, body RequestBody, options ...RequestOption) (*Response, error) {
	return DefaultClient.Delete(ctx, url, body, options...)
}
