package main

import (
	wb "github.com/alljames/weatherbit-go/pkg/api"
	"log"
)

// Parameters struct from the weatherbit package
// holds easily customisable fields to make an API request
// The weatherbit-go package takes care of the rest!
func populaterequestparameters(apikey string) wb.Parameters {

	P := wb.Parameters{}

	// EITHER specify P.city OR specify lat & lon

	// Steam Crane Pub, Bristol
	P.Lat = 51.4415
	P.Lon = -2.6017
	P.Temporality = "current" // "current", "history", "forecast"
	P.Apikey = apikey

	// FOR USE WITH HISTORY OR FORECAST QUERIES
	P.Granularity = "hourly"   // specify "hourly" OR "daily"
	P.StartDate = "2018-08-21" // [YYYY-MM-DD OR YYYY-MM-DD:HH]
	P.EndDate = "2018-08-22"   // [YYYY-MM-DD OR YYYY-MM-DD:HH]

	return P
}

func main() {

	log.Println("Starting")
	apikey, _ := wb.Readapikey("api_key.txt")
	p := populaterequestparameters(apikey)
	observation := wb.GetResponse(p)
	wb.Prettyprint(observation)
	log.Println("GHI:", observation.Data[0].Ghi)

}
