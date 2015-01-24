package godeckbrew

type Card struct {
	Artist       string   `json:"artist"`
	Cmc          int      `json:"cmc"`
	Colors       []string `json:"colors"`
	Flavor       string   `json:"flavor"`
	ImageName    string   `json:"imageName"`
	Layout       string   `json:"layout"`
	ManaCost     string   `json:"manaCost"`
	Multiverseid int      `json:"multiverseid"`
	Name         string   `json:"name"`
	Number       string   `json:"number"`
	Power        string   `json:"power"`
	Rarity       string   `json:"rarity"`
	Subtypes     []string `json:"subtypes"`
	Supertypes   []string `json:"supertypes"`
	Text         string   `json:"text"`
	Toughness    string   `json:"toughness"`
	Type         string   `json:"type"`
	Types        []string `json:"types"`
}
