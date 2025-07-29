package scryfall

import (
	"testing"
)

func TestNewClientUserAgent(t *testing.T) {
	clientTests := []struct {
		name              string
		clientOptions     []ClientOption
		expectedUserAgent string
	}{
		{name: "default user agent", clientOptions: nil, expectedUserAgent: "lorehold"},
		{name: "custom user agent", clientOptions: []ClientOption{WithUserAgent("test")}, expectedUserAgent: "test"},
	}

	for _, tt := range clientTests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.clientOptions...)
			got := client.userAgent
			if got != tt.expectedUserAgent {
				t.Errorf("got %s want %s", got, tt.expectedUserAgent)
			}
		})
	}
}

func TestNewBaseUri(t *testing.T) {
	clientTests := []struct {
		name            string
		clientOptions   []ClientOption
		expectedBaseUri string
	}{
		{name: "default base uri", clientOptions: nil, expectedBaseUri: "https://api.scryfall.com"},
		{name: "custom base uri", clientOptions: []ClientOption{WithBaseUri("http://test")}, expectedBaseUri: "http://test"},
	}

	for _, tt := range clientTests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.clientOptions...)
			got := client.baseUri
			if got != tt.expectedBaseUri {
				t.Errorf("got %s want %s", got, tt.expectedBaseUri)
			}
		})
	}
}
