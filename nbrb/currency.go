package nbrb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Currency struct.
// Ex. (Name: USD Rate: 2.50 Scale: 1)
type Currency struct {
	Name  string  `json:"Cur_Abbreviation"`
	Rate  float32 `json:"Cur_OfficialRate"`
	Scale int32   `json:"Cur_Scale"`
}

// FromBYN is convert currency from BYN
func (c *Currency) FromBYN(value float32) float32 {
	return value / (c.Rate / float32(c.Scale))
}

// ToBYN is convert currency to BYN
func (c *Currency) ToBYN(value float32) float32 {
	return value * (c.Rate / float32(c.Scale))
}

// Get currency rates from nbrb.by
func Get() *[]Currency {
	resp, err := http.Get("https://www.nbrb.by/api/exrates/rates/?periodicity=0")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var currencies []Currency

	if err := json.Unmarshal(body, &currencies); err != nil {
		log.Fatalln(err)
	}

	return &currencies
}

// PrintNeeded is print needed currencies
func PrintNeeded(currencies *[]Currency, neededCurrencies []string) {
	for _, v := range *currencies {
		if contains(neededCurrencies, v.Name) {
			fmt.Printf("%s %.4f\n", v.Name, v.Rate/float32(v.Scale))
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
