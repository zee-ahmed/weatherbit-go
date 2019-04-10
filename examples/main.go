package main

import (
	"fmt"
	wb "github.com/alljames/weatherbit-go/pkg"
	"os"
)

func main() {

	wb.SetEnvVariable("api_key.txt")
	// OR
	// wb.SetEnvVariable("0123456789abcdef")

	// Parameters struct from the weatherbit package
	// holds easily customisable fields to make an API request
	// The weatherbit-go package takes care of the rest!
	params := wb.Parameters{
		Lat:         51.4415, // Steam Crane Pub, Bristol
		Lon:         -2.6017,
		Temporality: "history", // "current", "history", "forecast"
		Apikey:      os.Getenv("WBITKEY"),

		// required for "history" or "forecast"  queries
		Granularity: "hourly",

		// required only for "history" queries
		StartDate: "2019-04-08", // [YYYY-MM-DD OR YYYY-MM-DD:HH]
		EndDate:   "2019-04-09", // [YYYY-MM-DD OR YYYY-MM-DD:HH]
	}

	fmt.Println(wb.GetResponse(params))

}
