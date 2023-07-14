package main

import (
	"datadazed.com/nasa-api-puller/apireader"
	"flag"
	"fmt"
)

func main() {
	var api string
	flag.StringVar(&api, "endpoint", "apod", "The API Endpoint you are wanting to pull")
	flag.Parse()

	fmt.Println(api)
	apiResponse := apireader.ValidateAPI(apireader.Endpoint(api))
	apireader.ReadAPI(apireader.Endpoint(apiResponse))
}
