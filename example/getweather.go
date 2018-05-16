package main

// import (
// 	wb "gitlab.com/aj16249/weatherbit-go/weatherbit"
// 	"time"
// )

// // If the function below causes a panic, one cause is a
// // 403 Forbidden status encountered when running in an environment
// // with limited permissions
// // fmt.Printf("%5.3f", results[0].Data[0].Dhi)
// func populatenowcast(n Nowcast, wbr wb.WbResponse) wb.Nowcast {
// 	timenow := time.Now().Unix()

// 	populatednowcast := wb.Nowcast{
// 		city:           wbr.Data[0].CityName,
// 		lat:            n.lat,
// 		lon:            n.lon,
// 		dhi:            wbr.Data[0].Dhi,
// 		windspeed:      wbr.Data[0].WindSpeed,
// 		cloudcover:     wbr.Data[0].CloudCover,
// 		temperature:    wbr.Data[0].Temperature,
// 		observationage: (timenow - int64(wbr.Data[0].LastObservationTimeStamp)) / 60,
// 	}

// 	return populatednowcast
// }
