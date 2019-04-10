package weatherbit

// TODO: import "net/url" would be a more idiomatic way to take care of generating HTTP GET requests

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func buildRequestURL(p Parameters) string {

	sb := strings.Builder{}
	sb.WriteString(BaseURL)

	if p.Temporality == "current" {
		sb.WriteString(p.Temporality + "?")
	} else if p.Temporality == "forecast" || p.Temporality == "history" {
		sb.WriteString(p.Temporality + "/" + p.Granularity + "?")
	}

	sb.WriteString(addlocation(p))
	sb.WriteString(addtimeframe(p))
	sb.WriteString("&key=" + os.Getenv("WBITKEY"))

	return sb.String()
}

func addlocation(p Parameters) string {

	if len(p.City) == 0 {
		return "lat=" + fmt.Sprint(p.Lat) + "&" + "lon=" + fmt.Sprint(p.Lon)
	}

	if p.Cityid != 0.0 { // Go initializes empty fields to zero values. p.cityid is a float64
		return "cityid=" + string(int(p.Cityid))
	}

	return "city=" + p.City
}

func addtimeframe(p Parameters) string {

	switch p.Temporality {
	case "current":
		return ""
	case "forecast":
		return ""
	case "history":
		return "&" + "start_date=" + p.StartDate + "&" + "end_date=" + p.EndDate

	default:
		log.Panicln("Temporality parameter not specified or not recognised. Must be \"current\", \"forecast\", or \"history\"")
	}

	// TODO: regexp to check that date is in correct format: YYYY-MM-DD for daily, YYYY-MM-DD:HH for hourly
	// Could even take current time and check that the dates given *are* in fact in the past / in the future

	return "ERROR - this code should be unreachable"
}
