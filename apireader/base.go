package apireader

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

type Endpoint string

var apiEndpoint = map[Endpoint]string{
	"apod": "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY",
}

func ValidateAPI(e Endpoint) string {
	// validate that the api endpoint exists
	endpoint, ok := apiEndpoint[e]
	if !ok {
		return fmt.Sprintf("Unsuppoprted api endpoint: %q", e)
	}
	return endpoint
}

func ReadAPI(url string) []byte {
	// Get the response from the API Call
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	// Defer closing the conncetion to the API
	defer response.Body.Close()

	// Validate Response Code
	responseStatus := response.StatusCode
	if responseStatus != 200 {
		fmt.Printf("Expected repsonse code 200, got: %v\n", responseStatus)
	}

	// Parse Response Data
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
