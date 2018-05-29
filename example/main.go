package main

import (
	"fmt"
	wb "github.com/alljames/weatherbit-go/weatherbit"
	"github.com/wcharczuk/go-chart"
	"net/http"
)

// Parameters struct from the weatherbit package
// holds easily customisable fields to make an API request
// The weatherbit-go package takes care of the rest!
func populaterequestparameters(apikey string) wb.Parameters {

	P := wb.Parameters{}

	// Steam Crane Pub, Bristol
	// P.City = "Johannesburg" // EITHER specify city OR specify lat & lon
	P.Lat = 51.4415
	P.Lon = -2.6017
	P.Temporality = "history" // "current", "history", "forecast"
	P.Apikey = apikey

	// FOR USE WITH HISTORY OR FORECAST QUERIES
	P.Granularity = "daily"    // specify "hourly" OR "daily"
	P.StartDate = "2018-05-20" // [YYYY-MM-DD OR YYYY-MM-DD:HH]
	P.EndDate = "2018-05-23"   // [YYYY-MM-DD OR YYYY-MM-DD:HH]

	return P
}

func main() {

	fmt.Println("Open browser to http://0.0.0.0:8080")
	http.HandleFunc("/", drawChart)
	http.ListenAndServe(":8080", nil)

}

func drawChart(resp http.ResponseWriter, req *http.Request) {

	apikey, wberr := wb.Readapikey("api_key.txt")
	if wberr != nil {
		panic("err")
	}
	p := populaterequestparameters(apikey)
	result := wb.GetResponse(p)
	resultlength := len(result.Data)

	xSlice := make([]float64, resultlength)
	ySlice := make([]float64, resultlength)

	for i := 0; i < resultlength; i++ {
		xSlice[i] = result.Data[i].LastObservationTimeStamp
		ySlice[i] = result.Data[i].Ghi
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true, //enables / displays the x-axis
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: xSlice,
				YValues: ySlice,
			},
		},
	}

	resp.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, resp)
}
