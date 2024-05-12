package main

import (
	"flag"
	"fmt"
	"bufio"
	"os"
	"log"
	"sync"

	"github.com/DEMON1A/mapz/pkg/validate"
)

func main() {
	// Setup the command line arguments
	var url string
	var file string
	var verbose bool
	var workers int
	var fast bool
	flag.StringVar(&url, "url", "", "The URL you want to validate the map existance for")
	flag.StringVar(&file, "file", "", "Path to your file that contains JavaScript URLs")
	flag.BoolVar(&verbose, "verbose", false, "Verbose mode to show additional information regarding the output")
	flag.IntVar(&workers, "workers", 4, "Number of workers to use for concurency")
	flag.BoolVar(&fast, "fast", false, "Perform a fast scan by disabling a validate rule, but it may result in some false positives")

	// Parse the flags
	flag.Parse()

	// Handle single urls in-case the user used a URL
	if url != "" {
		map_exists := validate.ValidateUrl(url, verbose, fast)
		if map_exists {
			fmt.Printf("%s.map\n", url)
		} else {
			if verbose {
				fmt.Printf("Didn't find a map file for %s\n", url)
			}
		}
		return
	}

	// Handle files
	if file != "" {
		// Open the file for reading
		file, err := os.Open(file)
		if err != nil {
			log.Fatalf("Error opening file: %s", err)
		}
		defer file.Close()

		// Create a new scanner to read from the file
		scanner := bufio.NewScanner(file)

		// Create a wait group to wait for all goroutines to finish
		var wg sync.WaitGroup

		// Channel to communicate results from goroutines
		resultCh := make(chan string)

		// Start a fixed number of worker goroutines
		numWorkers := workers // Adjust based on your system's capabilities
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for scanner.Scan() {
					line := string(scanner.Text())
					mapExists := validate.ValidateUrl(line, verbose, fast)
					if mapExists {
						resultCh <- line
					}
				}
			}()
		}

		// Start a goroutine to collect results from the result channel
		go func() {
			wg.Wait() // Wait for all worker goroutines to finish
			close(resultCh) // Close the result channel to signal completion
		}()

		// Print results received from the result channel
		for line := range resultCh {
			fmt.Printf("%s.map\n", line)
		}

		// Check for any errors during scanning
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error scanning file: %s", err)
		}
	}
}

// Mapz is given a URL like this https://example.com/static/main.js
// mapz is supposed to validate the existance of javascript map files
// sending an HTTP request to https://example.com/static/main.js.map
