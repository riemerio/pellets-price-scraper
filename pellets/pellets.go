package pellets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

var URL = "https://1heiz-pellets.de/x_preis.php"

func FetchPriceForLoosePellets(amountInKg int) PelletPrice {

	data := makeFormRequestForLoosePellets(amountInKg)

	resp, err := http.PostForm(URL, data)

	if err != nil {
		log.Fatal(err)
	}

	var res PelletPrice
	fmt.Println("loose pellets:", amountInKg)

	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func FetchPriceForSackedPellets(amountInPalettes int) PelletPrice {

	data := makeFormRequestForSackedPelletes(amountInPalettes)

	resp, err := http.PostForm(URL, data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sacked pellets:", amountInPalettes)
	var res PelletPrice

	json.NewDecoder(resp.Body).Decode(&res)

	return res
}

func makeFormRequestForSackedPelletes(amount int) url.Values {
	return makeFormRequest(amount, Sacked, UnitPallete)
}

func makeFormRequestForLoosePellets(amount int) url.Values {
	return makeFormRequest(amount, Loose, UnitKG)
}

func makeFormRequest(amount int, product string, unit string) url.Values {
	const ENV_NAME_POSTCODE = "SCRAPING_POSTCODE"
	postcode := os.Getenv(ENV_NAME_POSTCODE)

	if len([]rune(postcode)) != 5 {
		log.Fatalf("please provide a valid postcode as enviroment variable named: %s", ENV_NAME_POSTCODE)
	}

	return url.Values{
		"Produkt":              {product},
		"Einheit":              {unit},
		"Transport":            {"Lieferung"},
		"formURL":              {"%2Fpreise_holzpellets.php"},
		"Mehrwertsteuerfaktor": {"1.07"},
		"device":               {"c"},
		"Bundesland":           {"SN"},
		"Land":                 {"DE"},
		"plz":                  {postcode},
		"menge":                {strconv.Itoa(amount)},
		"lieferstellen":        {"1"},
	}
}

const (
	Loose       = "P101"
	Sacked      = "P104"
	UnitKG      = "kg"
	UnitPallete = "Pal."
)
