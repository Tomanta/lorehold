package scryfall

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	apiRoot        = "https://api.scryfall.com"
	userAgent      = "lorehold"
	acceptHeader   = "application/json;q=0.9,*/*;q=0.8"
	searchEndpoint = "/cards/search"
)

func CardSearch(query string) (*SearchResult, error) {
	httpClient := http.Client{Timeout: 10 * time.Second}

	url := apiRoot + searchEndpoint + "?q=" + query
	request, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", userAgent)
	request.Header.Set("Accept", acceptHeader)
	response, err := httpClient.Do(request)
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
