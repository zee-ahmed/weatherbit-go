package weatherbit

import (
	"github.com/wcharczuk/go-chart"
	"net/http"
)

// Swap out with main function in example/main.go
// func main() {

// 	fmt.Println("Open browser to http://0.0.0.0:8080")
// 	http.HandleFunc("/", drawChart)
// 	http.ListenAndServe(":8080", nil)

// }

func drawChart(resp http.ResponseWriter, req *http.Request) {

	apikey, wberr := Readapikey("api_key.txt")
	if wberr != nil {
		panic("err")
	}
	// p := Populaterequestparameters(apikey)
	p := Parameters{Apikey: apikey}
	result := GetResponse(p)
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
