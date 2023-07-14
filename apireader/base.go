package apireader

import "fmt"

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
	fmt.Println(endpoint)
	return endpoint
}

func ReadAPI(e Endpoint) string {
	return "works"
}
