package weatherbit

// TODO: import "net/url" would be a more idiomatic way to take care of generating HTTP GET requests

import (
	"fmt"
	"log"
)

// parameters struct holds relevant information for each request
type parameters struct {
	url         string
	apikey      string
	temporality string
	granularity string // "current", "forecast", "history"
	lat, lon    float64
	city        string
	state       string // "NC" or "North+Carolina"
	country     string // "US"
	cityid      float64
	units       string // "M" (Metric), "S" (Scientific), "I" (Imperial) - Default "M"
	marine      string // "f" (exclude offshore observations), "t" (include offshore observations) - Default "f"
	startDate   string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	endDate     string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
}

// Parameters - an exported struct which holds relevant information for each request
type Parameters struct {
	URL         string
	Apikey      string
	Temporality string
	Granularity string // "current", "forecast", "history"
	Lat, Lon    float64
	City        string
	State       string // "NC" or "North+Carolina"
	Country     string // "US"
	Cityid      float64
	Units       string // "M" (Metric), "S" (Scientific), "I" (Imperial) - Default "M"
	Marine      string // "f" (exclude offshore observations), "t" (include offshore observations) - Default "f"
	StartDate   string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	EndDate     string // YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
}

// createimportedparametersstruct allows a function to use the Parameters struct
// and map the data contained in it to an imported struct of type parameters
func createimportedparametersstruct(P Parameters) parameters {

	p := parameters{}

	p.apikey = P.Apikey
	p.temporality = P.Temporality
	p.granularity = P.Granularity
	p.lat = P.Lat
	p.lon = P.Lon
	p.city = P.City
	p.state = P.State
	p.country = P.Country
	p.startDate = P.StartDate
	p.endDate = P.EndDate

	p.cityid = P.Cityid
	p.units = P.Units
	p.marine = P.Marine

	return p
}

func buildRequestURL(p parameters) string {

	p.url = BaseURL + p.temporality + "/" + p.granularity
	p.url = addlocationtogetrequest(p)
	p.url = addtemoralityrogetrequest(p)
	p.url += "&" + "key=" + p.apikey

	return p.url
}

func addlocationtogetrequest(p parameters) string {

	p.url += "?"

	if len(p.city) == 0 {
		p.url += "lat=" + fmt.Sprint(p.lat) + "&" + "lon=" + fmt.Sprint(p.lon)
		return p.url
	}

	if p.cityid != 0.0 { // Go initiates empty fields to zero values. p.cityid is a float64
		p.url += "cityid=" + string(int(p.cityid))
		return p.url
	}

	p.url += "city=" + p.city
	return p.url
}

func addtemoralityrogetrequest(p parameters) string {

	if p.temporality == "current" || p.temporality == "forecast" {
		return p.url
	}

	if len(p.temporality) == 0 {
		log.Panicln("Temporality parameter not specified or not recognised. Must be \"current\", \"forecast\", or \"history\"")
	}

	// TODO: regexp to check that date is in correct format: YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	// Could even take current time and check that the dates given *are* in fact in the past / in the future

	if p.temporality == "history" {
		p.url += "&" + "start_date=" + p.startDate + "&" + "end_date=" + p.endDate
		return p.url
	}

	log.Panicln("Temporality parameter not specified or not recognised. Must be \"current\", \"forecast\", or \"history\"")
	return "ERROR"
}

// TODO: write and include an additional function that includes parameters metric/imperial/marine
