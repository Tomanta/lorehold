package scryfall

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

const (
	apiRoot        = "https://api.scryfall.com"
	userAgent      = "lorehold"
	acceptHeader   = "application/json;q=0.9,*/*;q=0.8"
	searchEndpoint = "/cards/search"
)

func CardSearch(apiClient *Client, query string) (*SearchResult, error) {
	url := apiRoot + searchEndpoint + "?q=" + query
	request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", apiClient.userAgent)
	request.Header.Set("Accept", apiClient.accepts)
	response, err := apiClient.Do(request)
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
