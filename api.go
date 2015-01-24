package godeckbrew

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
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

type Cents int

func (p Price) Cents() (Cents, error) {
	r := regexp.MustCompile(`\$(\d*?)\.(\d\d)`)
	matches := r.FindAllStringSubmatch(string(p), -1)
	if len(matches) != 1 || len(matches[0]) != 3 {
		return -1, fmt.Errorf("Could not parse price: %s", p)
	}
	mantissa, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return -1, fmt.Errorf("Could not parse price: %s", p)
	}
	decimal, err := strconv.Atoi(matches[0][2])
	if err != nil {
		return -1, fmt.Errorf("Could not parse price: %s", p)
	}

	return Cents(100*mantissa + decimal), nil

}

func (c *Card) Price() (p Cents, err error) {
	//TODO use set
	price, err := ChannelFireballPrice(c.Name, "")
	if err != nil {
		return -1, err
	}
	return price.Cents()
}

func Setlist(set string) (cards []*Card, err error) {
	const endpoint = "/mtg/cards"
	r := regexp.MustCompile(`<(.*?)>; rel="next"`)

	v := url.Values{}
	v.Set("set", set)
	page := 0
	for {
		u := baseUrl + endpoint + "?" + v.Encode()
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		bts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		log.Printf("Headers are %s", resp.Header.Get("Link"))
		linkHeader := resp.Header.Get("Link")

		var result []*Card
		err = json.Unmarshal(bts, &result)
		if err != nil {
			return nil, err
		}

		if nextMatch := r.FindStringSubmatch(linkHeader); len(nextMatch) == 2 {
			cards = append(cards, result...)
			page++
			v.Set("page", strconv.Itoa(page))
		} else {
			break
		}
	}

	return
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
