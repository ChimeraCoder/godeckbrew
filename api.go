package godeckbrew


//go:generate gojson -o card.go -input json/mtg/card.json -name "Card" -pkg "godeckbrew"

//go:generate gojson -o set.go -input json/mtg/set.json -name "Set" -pkg "godeckbrew"
