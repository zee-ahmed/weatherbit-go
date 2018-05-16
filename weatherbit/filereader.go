package weatherbit

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Readapikey takes API key file location and returns its contents as a string
func Readapikey(apikeyfile string) (string, error) {
	apikey, err := ioutil.ReadFile(apikeyfile)
	if err != nil {
		// TODO: read directory from which program is being run and print out the location?
		fmt.Println("Please acquire API key from Weatherbit.io and save in file helioscheduler/api_key.txt")
		log.Fatal(err)
	}

	apikeyread := strings.TrimSpace(string(apikey))
	if apikeyread == "" {
		log.Fatalln(apikeyfile + " did not contain API key")
	}

	log.Printf("API Key has been read from file \"%s\"\n", apikeyfile)
	return apikeyread, err
}
