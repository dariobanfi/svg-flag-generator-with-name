package getflag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pariz/gountries"
)

var ASSET_URL = "https://raw.githubusercontent.com/lipis/flag-icon-css/master/flags/4x3/"

var COUNTRY_NAME_RECT = `<svg width="640" height="480" x="140" y="340"> <g transform="translate(50,50)"> <rect rx="3" ry="3" width="260" height="70" stroke="black" fill="white" stroke-width="3"/> <svg width="260px" height="70px"> <text font-family="'Helvetica'" font-size="1.3em" x="50%" y="50%" alignment-baseline="middle" text-anchor="middle">country_name</text></svg></g></svg>`

func GetFlag(w http.ResponseWriter, r *http.Request) {
	flagid := strings.TrimPrefix(r.URL.Path, "/")

	if flagid == "" {
		fmt.Fprint(w, "Please insert a valid ISO 3166-1-alpha-2 code")
		return
	}

	// Gets country name from code
	query := gountries.New()
	country, _ := query.FindCountryByAlpha(flagid)
	countryname := country.Name.Common
	country_name_rect := strings.Replace(COUNTRY_NAME_RECT, "country_name", countryname, 1)

	resp, err := http.Get(ASSET_URL + flagid + ".svg")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	svg_code := string(body)
	svg_end_index := strings.Index(svg_code, "</svg>")
	svg_code = svg_code[:svg_end_index] + country_name_rect + svg_code[svg_end_index:]

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, `
<!DOCTYPE>
<html>
  <body>
`+svg_code+
		`</body>
</html>
  `)
}
