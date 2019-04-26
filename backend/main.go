package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

var (
	host = flag.String("host", "0.0.0.0:9999", "host to serve web")
)

type query struct {
	Country   string `json:"country"`
	Product   string `json:"product"`
	Revenue   string `json:"revenue"`
	Threshold string `json:"threshold"`
}

type answer struct {
	StandardRate string            `json:"stadartRate"`
	Categories   map[string]string `json:"categories"`
}

type Query interface {
	fetchData()
}

var eeJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/98c6a5ee-04e0-4608-9a26-24d1a91f5ae8.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/f6b62d72-a14b-4cec-b58b-1ae50ae290c9.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/4a02e91b-9f9b-4f7b-a21b-8af3fb5558ba.json",
}

var fiJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/a19cb68c-0d12-48f7-87f2-7cba255385ff.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/6b344392-cba4-4a34-a391-ebbb66fbf852.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/fe151deb-9e03-4739-af3b-afc06dd7219e.json",
}

var deJsonServices = []string{
	"http://ec.europa.eu/taxation_customs/tedb/api/search/8bf428cb-9152-46a3-b66a-3a922a4115cb.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/85e5e148-1e7a-47e9-b90e-022ede929c84.json",
	"http://ec.europa.eu/taxation_customs/tedb/api/search/7d5969a0-0fcc-4eb8-a0e8-cb919f578099.json",
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
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
		// output, err := json.Marshal(q)
		// if err != nil {
		// 	http.Error(w, err.Error(), 500)
		// 	return
		// }
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(q.fetchData()))
	})
	log.Fatal(http.ListenAndServe(*host, nil))
}

func (q query) fetchData() string {
	var a answer
	result := ""
	url := make([]string, 3)
	switch q.Country {
	case "EE":
		copy(url, eeJsonServices)
	case "FI":
		copy(url, fiJsonServices)
	case "DE":
		copy(url, deJsonServices)
	default:
		return result
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
			a.StandardRate = gjson.Get(bodyString, "results.0.standardRate").String()
			categoryList := gjson.Get(bodyString, "results.0.reducedRates.#.category")
			for _, name := range categoryList.Array() {
				println(gjson.Get(categoryList, "name").String())
			}
		}
		// decoder := json.NewDecoder(res.Body)
		// var data Tracks
		// err = decoder.Decode(&data)
		// if err != nil {
		// 	panic(err)
		// }

		res.Body.Close()
	}

	return result
}
