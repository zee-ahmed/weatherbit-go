package weatherbit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Example URL
// https://api.weatherbit.io/v2.0/current?lat=LATITUDE&lon=LONGITUDE&key=APIKEY
func syncHTTPGets(p parameters) WbResponse {

	url := buildRequestURL(p)
	log.Println(url)

	wbClient := http.Client{
		Timeout: time.Second * ReqTimeout,
	}

	req, reqErr := http.NewRequest(http.MethodGet, url, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}

	// Golang encrypts and decrypts HTTP requests (gzip) by default
	req.Header.Set("User-Agent", "application/json, */*")
	resp, getErr := wbClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var wbresp WbResponse

	jsonErr := json.Unmarshal(body, &wbresp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return wbresp
}
