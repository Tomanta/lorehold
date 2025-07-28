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
)

type Client struct {
	ApiClient *http.Client
	userAgent string
	accepts   string
	rateLimit time.Duration
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: timeoutDefault,
	}

	apiClient := &Client{
		ApiClient: httpClient,
		userAgent: userAgentDefault,
		accepts:   acceptsDefault,
		rateLimit: rateLimitDefault,
	}
	return apiClient
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.ApiClient.Do(req)
	return resp, err
}
