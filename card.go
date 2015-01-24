package godeckbrew

type Card struct {
	Cmc      int      `json:"cmc"`
	Colors   []string `json:"colors"`
	Cost     string   `json:"cost"`
	Editions []struct {
		Artist       string `json:"artist"`
		Flavor       string `json:"flavor"`
		ImageURL     string `json:"image_url"`
		Layout       string `json:"layout"`
		MultiverseID int    `json:"multiverse_id"`
		Number       string `json:"number"`
		Price        struct {
			High   int `json:"high"`
			Low    int `json:"low"`
			Median int `json:"median"`
		} `json:"price"`
		Rarity   string `json:"rarity"`
		Set      string `json:"set"`
		SetID    string `json:"set_id"`
		SetURL   string `json:"set_url"`
		StoreURL string `json:"store_url"`
		URL      string `json:"url"`
	} `json:"editions"`
	Formats struct {
		Commander string `json:"commander"`
		Legacy    string `json:"legacy"`
		Vintage   string `json:"vintage"`
	} `json:"formats"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	StoreURL string   `json:"store_url"`
	Text     string   `json:"text"`
	Types    []string `json:"types"`
	URL      string   `json:"url"`
}
