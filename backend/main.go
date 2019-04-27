package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
)

var (
	host = flag.String("host", "0.0.0.0:9999", "host to serve web")
)

type query struct {
	Country string `json:"country"`
}

type countryData struct {
	StandardRate string         `json:"standardRate"`
	Threshold    string         `json:"threshold"`
	Categories   []categoryData `json:"categories"`
}

type categoryData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Comments    string `json:"comments"`
	ReducedRate string `json:"reducedRate"`
}
type Query interface {
	fetchData()
}

var eeJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/98c6a5ee-04e0-4608-9a26-24d1a91f5ae8.json",
}

var fiJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/a19cb68c-0d12-48f7-87f2-7cba255385ff.json",
}

var deJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/8bf428cb-9152-46a3-b66a-3a922a4115cb.json",
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		var q query
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&q)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		answer, err := json.Marshal(q.fetchData())
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(answer)
	}).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(*host, r))
}

func (q query) fetchData() countryData {
	var country countryData

	url := make([]string, 1)

	switch q.Country {
	case "EE":
		copy(url, eeJsonServices)
	case "FI":
		copy(url, fiJsonServices)
	case "DE":
		copy(url, deJsonServices)
	default:
		return country
	}

	for _, u := range url {
		res, err := http.Get(u)
		if err != nil {
			panic(err)
		}
		if res.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			country.StandardRate = gjson.Get(bodyString, `results.0.standardRate`).String()
			country.Threshold = gjson.Get(bodyString, `results.0.specialSchemeDistanceSellingThreshold`).String()
			categoryList := gjson.Get(bodyString, "results.0.reducedRates")
			for _, name := range categoryList.Array() {
				var category categoryData
				category.Name = gjson.Get(name.String(), "category.name").String()
				category.ReducedRate = gjson.Get(name.String(), "reducedRate").String()
				category.Comments = gjson.Get(name.String(), "comments").String()
				category.Description = gjson.Get(name.String(), "category.description").String()
				country.Categories = append(country.Categories, category)
			}
		}
		res.Body.Close()
	}

	return country
}
