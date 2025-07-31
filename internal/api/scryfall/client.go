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
	apiClient   *http.Client
	userAgent   string
	accepts     string
	rateLimit   time.Duration
	baseUri     string
	lastRequest time.Time
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

func WithTimeout(timeout time.Duration) ClientOption {
	return func(o *clientOptions) {
		o.rateLimit = timeout
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
	// Not optimal algorithm; if we add the limit duration to
	// the last request time and that is before NOW,
	// sleep for the rate limit. Latency is minimal for a
	// desktop app.
	if time.Now().Before(c.lastRequest.Add(c.rateLimit)) {
		time.Sleep(c.rateLimit)
	}
	resp, err := c.apiClient.Do(req)
	c.lastRequest = time.Now()
	return resp, err
}
