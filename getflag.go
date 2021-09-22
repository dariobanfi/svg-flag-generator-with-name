package getflag

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/pariz/gountries"
)

var AssetUrl = "https://raw.githubusercontent.com/lipis/flag-icon-css/master/flags/4x3/"

var CountryNameRect = `<svg width="640" height="480" x="140" y="340"> <g transform="translate(50,50)"> <rect rx="3" ry="3" width="260" height="70" stroke="black" fill="white" stroke-width="3"/> <svg width="260px" height="70px"> <text font-family="'Helvetica'" font-size="1.3em" x="50%" y="50%" alignment-baseline="middle" text-anchor="middle">country_name</text></svg></g></svg>`

func GetFlag(w http.ResponseWriter, r *http.Request) {

	flagId := strings.TrimPrefix(r.URL.Path, "/")

	if flagId == "" {
		fmt.Fprint(w, "Seleziona un paese aggiungendo il codice  ISO 3166-1 nell'URL (e.g. /it)")
		return
	}

	// Gets country name from code
	query := gountries.New()
	country, _ := query.FindCountryByAlpha(flagId)
	countryName := getCountryNameForLocale(country, "ITA")
	countryNameRect := strings.Replace(CountryNameRect, "country_name", countryName, 1)

	// Downloads the flag svg
	resp, err := http.Get(AssetUrl + flagId + ".svg")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	svg_code := string(body)
	svg_end_index := strings.Index(svg_code, "</svg>")
	svg_code = svg_code[:svg_end_index] + countryNameRect + svg_code[svg_end_index:]

	// Builds list of all countries
	countries := query.FindAllCountries()
	var sb strings.Builder
	sb.WriteString("<ul>")
	for key, element := range countries {
		sb.WriteString("<li><a href=\"" + strings.ToLower(key) + "\">" + getCountryNameForLocale(element, "ITA") + "</a></li>\n")

	}
	sb.WriteString("</ul>")

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprint(w, `
<!DOCTYPE>
<html>
  <head>
    <script>
      function downloadSVG() {
        const svg = document.getElementById('svgflag').innerHTML;
        const blob = new Blob([svg.toString()]);
        const element = document.createElement("a");
        element.download = "`+flagId+`.svg";
        element.href = window.URL.createObjectURL(blob);
        element.click();
        element.remove();
      }
    </script>
  </head>
  <body>
  <h1>Bandiera del Paese: `+countryName+`</h1>
  <div style="display:flex; flex-direction: row; justify-content: space-evenly;">
    <div style="width: 50%; text-align: center;">
      <div id="svgflag" style="border: 4px solid black;">`+svg_code+`</div>
      <br /><br /><br />
      <button onclick="downloadSVG()"> DOWNLOAD </button>
    </div>
  </div>
  <div>
      Altre bandiere:`+sb.String()+
		`<div/>
  </body>
</html>
  `)
}

func getCountryNameForLocale(country gountries.Country, locale string) string {
	value := country.Translations[locale].Common
	if len(value) == 0 {
		return country.Name.Common
	}
	return value
}
