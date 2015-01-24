package godeckbrew

type Set struct {
	Block        string        `json:"block"`
	Booster      []interface{} `json:"booster"`
	Border       string        `json:"border"`
	Cards        []Card        `json:"cards"`
	Code         string        `json:"code"`
	GathererCode string        `json:"gathererCode"`
	Name         string        `json:"name"`
	OldCode      string        `json:"oldCode"`
	OnlineOnly   bool          `json:"onlineOnly"`
	ReleaseDate  string        `json:"releaseDate"`
	Type         string        `json:"type"`
}
