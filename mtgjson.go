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
const _CardIndexUrl = "http://mtgjson.com/json/"

// DownloadSet will download the JSON representation of cards if it does not exist
// "set" must be a three-letter string (e.g. KTK for Khans of Tarkir)
func DownloadSet(set string) error {

    if set == "" {
        set = _CardIndexFilename
    }
	filename := path.Join(app.UserData(), set + ".json")
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
		log.Printf("Fetching url %s", _CardIndexUrl + set + ".json")
		resp, err := http.Get(_CardIndexUrl + set + ".json")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		n, err := io.Copy(out, resp.Body)
		log.Printf("Wrote %d bytes", n)
	}
	return nil
}
