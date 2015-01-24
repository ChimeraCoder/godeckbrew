package godeckbrew

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//go:generate gojson -o card.go -input json/mtg/card.json -name "Card" -pkg "godeckbrew"

//go:generate gojson -o set.go -input json/mtg/set.json -name "Set" -pkg "godeckbrew"

const baseUrl = "https://api.deckbrew.com"

// GetCard implements the /mtg/cards/<id> endpoint
func GetCard(id string) (card Card, err error) {
	const endpoint = "/mtg/cards/"
	resp, err := http.Get(baseUrl + endpoint + id)
	if err != nil {
		return
	}
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bts, &card)
	return
}

type Price string

func (c *Card) Price() (p Price, err error) {
	//TODO use set
	price, err := ChannelFireballPrice(c.Name, "")
	if err != nil {
		return
	}
	return price, nil
}

// ChannelFireballPrice fetches the price from Channel Fireball
// The card should be the full name (e.g. "Dark Confidant") and the set
// (e.g. "ravnica") is optional
// If the set is not provided, this will return the price for the first set returned by CFB
func ChannelFireballPrice(card string, set string) (price Price, err error) {
	const u = "http://magictcgprices.appspot.com/api/cfb/price.json?"
	v := url.Values{}
	v.Set("cardname", card)
	v.Set("setname", set)
	resp, err := http.Get(u + v.Encode())
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result []string
	err = json.Unmarshal(bts, &result)
	if len(result) == 0 {
		return "", fmt.Errorf("No prices found")
	}

	// Currently CFB only returns a single price

	price = Price(result[0])

	return price, err
}
