package main

import (
	wb "github.com/alljames/weatherbit-go/pkg"
	"log"
)

func main() {

	log.Println("Starting")
	apikey, _ := wb.Readapikey("api_key.txt")

	// Parameters struct from the weatherbit package
	// holds easily customisable fields to make an API request
	// The weatherbit-go package takes care of the rest!
	p := wb.Parameters{
		Lat:         51.4415, // Steam Crane Pub, Bristol
		Lon:         -2.6017,
		Temporality: "history", // "current", "history", "forecast"
		Apikey:      apikey,

		// required for "history" or "forecast"  queries
		Granularity: "hourly",

		// required only for "history" queries
		StartDate: "2018-12-21", // [YYYY-MM-DD OR YYYY-MM-DD:HH]
		EndDate:   "2018-12-22", // [YYYY-MM-DD OR YYYY-MM-DD:HH]
	}

	observation := wb.GetResponse(p)
	wb.Prettyprint(observation)

}
