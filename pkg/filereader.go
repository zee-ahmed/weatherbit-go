package weatherbit

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Readapikey takes API key file location and returns its contents as a string
func SetEnvVariable(apiKeyFile string) (string, error) {

	// TODO: Environment variables should be used here

	var err error

	if len(os.Getenv("WBITKEY")) == 1 {
		return os.Getenv("WBITKEY"), err
	}

	// os.Getenv("WBITKEY")
	apikeyraw, err := ioutil.ReadFile(apiKeyFile)
	if err != nil {
		// TODO: read directory from which program is being run and print out the location?
		fmt.Println("Please acquire API key from Weatherbit.io and save in file " + apiKeyFile)
		log.Fatal(err)
	}

	apikey := strings.TrimSpace(string(apikeyraw))
	if apikey == "" {
		log.Fatalln(apiKeyFile + " did not contain API key")
	}

	log.Printf("API Key has been read from file \"%s\"\n", apiKeyFile)
	os.Setenv("WBITKEY", apikey)

	return apikey, err
}
