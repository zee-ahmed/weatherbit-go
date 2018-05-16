package weatherbit

import (
	"fmt"
	"math"
	"time"
)

const (
	degtorad float64 = 0.01745329252
)

// GetResponse is
func GetResponse(ep Parameters) WbResponse {
	p := createimportedparametersstruct(ep)
	response := syncHTTPGets(p)
	return response
}

// Prettyprint can be customised by the user to print the elements of
// their choosing within the WbResponse struct, but is mostly here to demonstrate
// how to access elements out of the WbResponse struct
func Prettyprint(wbr WbResponse) {

	fmt.Printf("City: %s\n", wbr.Data[0].CityName)
	count := len(wbr.Data)
	fmt.Printf("Number of results returned: %d\n", count)
	for i := 0; i < count; i++ {

		observationtimestamp := int64(wbr.Data[i].LastObservationTimeStamp)
		observationtime := time.Unix(observationtimestamp, (observationtimestamp / 1e9))

		dni := wbr.Data[i].Dni
		dhi := wbr.Data[i].Dhi
		solarelevationangle := wbr.Data[i].SolarElevationAngle // degrees
		solarelevationradians := solarelevationangle * degtorad
		ghi := (dni*math.Cos(solarelevationradians) + wbr.Data[i].Dhi)
		fmt.Printf("Observation time: %v\t (%v)\n", observationtime, (time.Until(observationtime)))
		fmt.Printf("Temperature: %f\t", wbr.Data[i].Temperature)
		fmt.Printf("CloudsHi: %f\n", wbr.Data[i].CloudsHi)
		fmt.Printf("GHI: %f\n", ghi)
		fmt.Printf("UV: %f\t DNI: %f\t DHI: %f\t", wbr.Data[i].UV, dni, dhi)
		fmt.Printf("SolarElevationAngle: %f\t", wbr.Data[i].SolarElevationAngle)
		fmt.Printf("SolarHourAngle: %f\t", wbr.Data[i].SolarHourAngle)
		// Global Horizontal (GHI) = Direct Normal (DNI) X cos(θ) + Diffuse Horizontal (DHI)
		// TODO: check degree / radians. use golang geo package to convert between the two
	}

}
