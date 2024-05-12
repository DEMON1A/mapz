package validate

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"bytes"
)

func ValidateUrl(url string, verbose bool, fast bool) bool {
	// Craft the .js.map url
	mapURL := fmt.Sprintf("%s.map", url)

	// Create a new GET request
	req, err := http.NewRequest("GET", mapURL, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	// Send the HTTP request using http.DefaultClient
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	// Print the status code of the response
	if verbose {
		fmt.Printf("%s did return a status code: %d\n", url, resp.StatusCode)
	}
	
	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return false
	}

	// Process response body in chunks to avoid loading large files into memory at once
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseBody := buf.String()

	if !fast {
		// Check for required substrings in the response body
		if strings.Contains(responseBody, "\"version\"") &&
			strings.Contains(responseBody, "\"sources\"") {
			return true
		}
	} else {
		return true
	}

	return false
}
