package getflag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetFlag(w http.ResponseWriter, r *http.Request) {
	flagid := strings.TrimPrefix(r.URL.Path, "/")
	asset_url := "https://raw.githubusercontent.com/lipis/flag-icon-css/master/flags/4x3/"

	if flagid == "" {
		fmt.Fprint(w, "Please insert a valid ISO 3166-1-alpha-2 code")
		return
	}

	resp, err := http.Get(asset_url + flagid + ".svg")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	// w.Header().Add("Content-Type", "image/svg")
	fmt.Fprint(w, sb)
}
