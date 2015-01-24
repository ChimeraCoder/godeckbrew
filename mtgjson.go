package godeckbrew

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/Wessie/appdirs"
)

var app = appdirs.New("magictcg", "chimeracoder", ".1")

const _CardIndexFilename = "Allsets-x.json"
const _CardIndexUrl = "http://mtgjson.com/json/AllSets-x.json"
orpusUrl

// DownloadCards will download the JSON representation of cards if it does not exist
func DownloadCards() error {
	filename := path.Join(app.UserData(), _CardIndexFilename)
	// Check if file already exists
	if _, err := os.Stat(filename); err != nil {

		log.Printf("Writing to filename %s", filename)

		err := os.MkdirAll(app.UserData(), os.ModePerm)
		if err != nil {
			return err
		}
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
			return err
		}
		defer out.Close()
		log.Printf("Fetching url %s", _CardIndexUrl)
		resp, err := http.Get(_CardIndexUrl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		n, err := io.Copy(out, resp.Body)
		log.Printf("Wrote %d bytes", n)
	}
	return nil
}
