package scryfall

import (
	"net/http"
	"time"
)

const (
	userAgentDefault string        = "lorehold"
	acceptsDefault   string        = "application/json;q=0.9,*/*;q=0.8"
	timeoutDefault   time.Duration = 10 * time.Second

	// Scryfall requests no more than 10 requests per second
	rateLimitDefault time.Duration = 100 * time.Millisecond

	// base URL
	baseUriDefault string = "https://api.scryfall.com"
)

type Client struct {
	apiClient *http.Client
	userAgent string
	accepts   string
	rateLimit time.Duration
	baseUri   string
}

type clientOptions struct {
	userAgent string
	accepts   string
	rateLimit time.Duration
	baseUri   string
}

type ClientOption func(*clientOptions)

func WithBaseUri(uri string) ClientOption {
	return func(o *clientOptions) {
		o.baseUri = uri
	}
}

func WithUserAgent(userAgent string) ClientOption {
	return func(o *clientOptions) {
		o.userAgent = userAgent
	}
}

func NewClient(options ...ClientOption) *Client {
	opts := &clientOptions{
		userAgent: userAgentDefault,
		accepts:   acceptsDefault,
		rateLimit: rateLimitDefault,
		baseUri:   baseUriDefault,
	}

	for _, option := range options {
		option(opts)
	}

	httpClient := &http.Client{
		Timeout: timeoutDefault,
	}

	apiClient := &Client{
		apiClient: httpClient,
		userAgent: opts.userAgent,
		accepts:   opts.accepts,
		rateLimit: opts.rateLimit,
		baseUri:   opts.baseUri,
	}
	return apiClient
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.apiClient.Do(req)
	return resp, err
}
