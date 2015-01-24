package godeckbrew

import (
	"math/rand"
	"strings"
)

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

// BoosterSize returns the number of cards in a booster pack for the set
func (s Set) BoosterSize() int {
	return len(s.Booster)
}

// NewBoosterPack returns a slice of 15 cards with the appropriate rarity distributions
func (s Set) NewBoosterPack() []*Card {
	// TODO support other booster distributions
	cards := make([]*Card, 15)

	mythicRoll := rand.Intn(8)
	if mythicRoll == 7 {
		// first slot is a mythic rare
		mythicRares := s.FilterRarity("mythicRare")
		randCard := mythicRares[rand.Intn(len(mythicRares))]
		cards[0] = &randCard
	} else {
		// first slot is a rare
		rares := s.FilterRarity("rare")
		randCard := rares[rand.Intn(len(rares))]

		cards[0] = &randCard
	}

	// TODO sample without replacement
	// Draw 3 uncommons
	uncommons := s.FilterRarity("uncommon")
	for i := 1; i < 4; i++ {
		randCard := uncommons[rand.Intn(len(uncommons))]

		cards[i] = &randCard
	}

	// Draw 11 commons
	commons := s.FilterRarity("common")
	for i := 4; i < 15; i++ {
		randCard := commons[rand.Intn(len(commons))]

		cards[i] = &randCard
	}
	return cards
}

func (s Set) FilterRarity(rarity string) []Card {
	result := []Card{}
	for _, card := range s.Cards {
		if strings.ToLower(card.Rarity) == rarity {
			result = append(result, card)
		}
	}
	return result
}
