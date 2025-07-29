package scryfall

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	searchEndpoint = "/cards/search"
)

func (c *Client) Search(query string) (*SearchResult, error) {
	url := c.baseUri + searchEndpoint + "?q=" + query
	request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", c.userAgent)
	request.Header.Set("Accept", c.accepts)
	response, err := c.do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResult SearchResult
	err = json.Unmarshal(data, &searchResult)
	if err != nil {
		return nil, err
	}

	return &searchResult, nil

}
