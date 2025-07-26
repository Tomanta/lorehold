package scryfall

const (
	apiRoot        = "https://api.scryfall.com"
	userAgent      = "lorehold"
	acceptHeader   = "application/json;q=0.9,*/*;q=0.8"
	searchEndpoint = "/cards/search"
)

func CardSearch(query string) string {
	return "Search results"
}
