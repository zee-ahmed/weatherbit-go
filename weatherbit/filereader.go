package weatherbit

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Readapikey takes API key file location and returns its contents as a string
func Readapikey(apiKeyFile string) (string, error) {
	apikey, err := ioutil.ReadFile(apiKeyFile)
	if err != nil {
		// TODO: read directory from which program is being run and print out the location?
		fmt.Println("Please acquire API key from Weatherbit.io and save in file helioscheduler/api_key.txt")
		log.Fatal(err)
	}

	apiKeyRead := strings.TrimSpace(string(apikey))
	if apiKeyRead == "" {
		log.Fatalln(apiKeyFile + " did not contain API key")
	}

	log.Printf("API Key has been read from file \"%s\"\n", apiKeyFile)
	return apiKeyRead, err
}
