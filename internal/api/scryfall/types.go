package scryfall

type AllPart struct {
	Object    string `json:"object"`
	ID        string `json:"id"`
	Component string `json:"component"`
	Name      string `json:"name"`
	TypeLine  string `json:"type_line"`
	URI       string `json:"uri"`
}

type CardFace struct {
	Object         string    `json:"object"`
	Name           string    `json:"name"`
	ManaCost       string    `json:"mana_cost"`
	TypeLine       string    `json:"type_line"`
	OracleText     string    `json:"oracle_text"`
	Colors         []string  `json:"colors"`
	Power          string    `json:"power"`
	Toughness      string    `json:"toughness"`
	FlavorText     string    `json:"flavor_text"`
	Artist         string    `json:"artist"`
	ArtistID       string    `json:"artist_id"`
	IllustrationID string    `json:"illustration_id"`
	ImageUris      ImageUris `json:"image_uris"`
	ColorIndicator []string  `json:"color_indicator,omitempty"`
}

type ImageUris struct {
	Small      string `json:"small"`
	Normal     string `json:"normal"`
	Large      string `json:"large"`
	Png        string `json:"png"`
	ArtCrop    string `json:"art_crop"`
	BorderCrop string `json:"border_crop"`
}

type Legalities struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Timeless        string `json:"timeless"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Oathbreaker     string `json:"oathbreaker"`
	Standardbrawl   string `json:"standardbrawl"`
	Brawl           string `json:"brawl"`
	Alchemy         string `json:"alchemy"`
	Paupercommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	Oldschool       string `json:"oldschool"`
	Premodern       string `json:"premodern"`
	Predh           string `json:"predh"`
}

type Preview struct {
	Source      string `json:"source"`
	SourceURI   string `json:"source_uri"`
	PreviewedAt string `json:"previewed_at"`
}

type Prices struct {
	Usd       string `json:"usd"`
	UsdFoil   string `json:"usd_foil"`
	UsdEtched any    `json:"usd_etched"`
	Eur       string `json:"eur"`
	EurFoil   string `json:"eur_foil"`
	Tix       string `json:"tix"`
}

type PurchaseUris struct {
	Tcgplayer   string `json:"tcgplayer"`
	Cardmarket  string `json:"cardmarket"`
	Cardhoarder string `json:"cardhoarder"`
}

type RelatedUris struct {
	Gatherer                  string `json:"gatherer"`
	TcgplayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TcgplayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	Edhrec                    string `json:"edhrec"`
}

type CardData struct {
	Object          string       `json:"object"`
	ID              string       `json:"id"`
	OracleID        string       `json:"oracle_id"`
	MultiverseIds   []int        `json:"multiverse_ids"`
	MtgoID          int          `json:"mtgo_id,omitempty"`
	TcgplayerID     int          `json:"tcgplayer_id,omitempty"`
	CardmarketID    int          `json:"cardmarket_id,omitempty"`
	Name            string       `json:"name"`
	Lang            string       `json:"lang"`
	ReleasedAt      string       `json:"released_at"`
	URI             string       `json:"uri"`
	ScryfallURI     string       `json:"scryfall_uri"`
	Layout          string       `json:"layout"`
	HighresImage    bool         `json:"highres_image"`
	ImageStatus     string       `json:"image_status"`
	ImageUris       ImageUris    `json:"image_uris,omitempty"`
	ManaCost        string       `json:"mana_cost,omitempty"`
	Cmc             float64      `json:"cmc"`
	TypeLine        string       `json:"type_line"`
	OracleText      string       `json:"oracle_text,omitempty"`
	Power           string       `json:"power,omitempty"`
	Toughness       string       `json:"toughness,omitempty"`
	Colors          []string     `json:"colors,omitempty"`
	ColorIndicator  []string     `json:"color_indicator,omitempty"`
	ColorIdentity   []string     `json:"color_identity"`
	Keywords        []any        `json:"keywords"`
	AllParts        []AllPart    `json:"all_parts,omitempty"`
	Legalities      Legalities   `json:"legalities"`
	Games           []string     `json:"games"`
	Reserved        bool         `json:"reserved"`
	GameChanger     bool         `json:"game_changer"`
	Foil            bool         `json:"foil"`
	Nonfoil         bool         `json:"nonfoil"`
	Finishes        []string     `json:"finishes"`
	Oversized       bool         `json:"oversized"`
	Promo           bool         `json:"promo"`
	Reprint         bool         `json:"reprint"`
	Variation       bool         `json:"variation"`
	SetID           string       `json:"set_id"`
	Set             string       `json:"set"`
	SetName         string       `json:"set_name"`
	SetType         string       `json:"set_type"`
	SetURI          string       `json:"set_uri"`
	SetSearchURI    string       `json:"set_search_uri"`
	ScryfallSetURI  string       `json:"scryfall_set_uri"`
	RulingsURI      string       `json:"rulings_uri"`
	PrintsSearchURI string       `json:"prints_search_uri"`
	CollectorNumber string       `json:"collector_number"`
	Digital         bool         `json:"digital"`
	Rarity          string       `json:"rarity"`
	CardBackID      string       `json:"card_back_id,omitempty"`
	Artist          string       `json:"artist"`
	ArtistIds       []string     `json:"artist_ids"`
	IllustrationID  string       `json:"illustration_id,omitempty"`
	BorderColor     string       `json:"border_color"`
	Frame           string       `json:"frame"`
	FrameEffects    []string     `json:"frame_effects,omitempty"`
	SecurityStamp   string       `json:"security_stamp,omitempty"`
	FullArt         bool         `json:"full_art"`
	Textless        bool         `json:"textless"`
	Booster         bool         `json:"booster"`
	StorySpotlight  bool         `json:"story_spotlight"`
	EdhrecRank      int          `json:"edhrec_rank,omitempty"`
	PennyRank       int          `json:"penny_rank,omitempty"`
	Preview         Preview      `json:"preview,omitempty"`
	Prices          Prices       `json:"prices"`
	RelatedUris     RelatedUris  `json:"related_uris,omitempty"`
	PurchaseUris    PurchaseUris `json:"purchase_uris,omitempty"`
	ArenaID         int          `json:"arena_id,omitempty"`
	CardFaces       []CardFace   `json:"card_faces,omitempty"`
	FlavorText      string       `json:"flavor_text,omitempty"`
	MtgoFoilID      int          `json:"mtgo_foil_id,omitempty"`
	PromoTypes      []string     `json:"promo_types,omitempty"`
	ProducedMana    []string     `json:"produced_mana,omitempty"`
	Watermark       string       `json:"watermark,omitempty"`
}

type SearchResult struct {
	Object     string     `json:"object"`
	TotalCards int        `json:"total_cards"`
	HasMore    bool       `json:"has_more"`
	NextPage   string     `json:"next_page"`
	Data       []CardData `json:"data"`
}
