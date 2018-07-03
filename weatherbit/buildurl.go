package weatherbit

// TODO: import "net/url" would be a more idiomatic way to take care of generating HTTP GET requests

import (
	"fmt"
	"log"
	"strings"
)

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

	sb := strings.Builder{}
	sb.WriteString(BaseURL)

	if p.temporality == "current" {
		sb.WriteString(p.temporality + "?")
	} else if p.temporality == "forecast" || p.temporality == "history" {
		sb.WriteString(p.temporality + "/" + p.granularity + "?")
	}

	sb.WriteString(addlocation(p))
	sb.WriteString(addtimeframe(p))
	sb.WriteString("&key=" + p.apikey)

	p.url = sb.String()

	return p.url
}

func addlocation(p parameters) string {

	if len(p.city) == 0 {
		return "lat=" + fmt.Sprint(p.lat) + "&" + "lon=" + fmt.Sprint(p.lon)
	}

	if p.cityid != 0.0 { // Go initiates empty fields to zero values. p.cityid is a float64
		return "cityid=" + string(int(p.cityid))
	}

	return "city=" + p.city
}

func addtimeframe(p parameters) string {

	switch p.temporality {
	case "current":
		return ""
	case "forecast":
		return ""
	case "history":
		return "&" + "start_date=" + p.startDate + "&" + "end_date=" + p.endDate

	default:
		log.Panicln("Temporality parameter not specified or not recognised. Must be \"current\", \"forecast\", or \"history\"")
	}

	// TODO: regexp to check that date is in correct format: YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	// Could even take current time and check that the dates given *are* in fact in the past / in the future

	return "ERROR - this code should be unreachable"
}
