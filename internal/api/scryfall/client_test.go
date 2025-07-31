package scryfall

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func createTestServer(path string, handler func(http.ResponseWriter, *http.Request), clientOptions ...ClientOption) (*Client, *httptest.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc(path, handler)
	ts := httptest.NewServer(mux)

	mergedClientOptions := []ClientOption{WithBaseUri(ts.URL), WithTimeout(time.Second * 0)}
	mergedClientOptions = append(mergedClientOptions, clientOptions...)
	client := NewClient(mergedClientOptions...)
	return client, ts, nil

}

func TestNewClientSendsUserAgent(t *testing.T) {
	tests := []struct {
		name              string
		clientOptions     []ClientOption
		expectedUserAgent string
	}{
		{
			name:              "default user agent",
			clientOptions:     nil,
			expectedUserAgent: userAgentDefault,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				userAgent := r.Header.Get("User-Agent")
				if userAgent != test.expectedUserAgent {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintln(w, `{"object": "error", "code":"bad_request", "status":400, "details": ""}`)
					return
				}

				fmt.Fprintln(w, `{"object":"list", "has_more": false, "data":[]}`)
			})

			client, ts, err := createTestServer("/cards/search", handler, test.clientOptions...)
			if err != nil {
				t.Fatalf("Error setting up test server: %v", err)
			}
			defer ts.Close()

			_, err = client.Search("test")
			if err != nil {
				t.Fatalf("Error validating user agent: %v", err)
			}
		})
	}

}
func TestNewClientUserAgent(t *testing.T) {
	clientTests := []struct {
		name              string
		clientOptions     []ClientOption
		expectedUserAgent string
	}{
		{name: "default user agent", clientOptions: nil, expectedUserAgent: userAgentDefault},
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
