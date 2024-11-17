package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	// Define command line flags for the input file and the site URL
	inputFile := flag.String("list", "", "Path to the input file")
	siteURL := flag.String("site", "", "URL to replace the parameter value with")
	flag.Parse()

	// Check if the input file flag is provided
	if *inputFile == "" {
		fmt.Println("Please provide an input file using the -list flag")
		return
	}

	// Check if the site URL flag is provided
	if *siteURL == "" {
		fmt.Println("Please provide a site URL using the -site flag")
		return
	}

	// Open the input file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a custom HTTP client with a redirect policy that does not follow redirects
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Replace the URL parameter value with the provided site URL
		modifiedLine := replaceURLParameter(line, *siteURL)
		// Make a web request to the modified URL
		resp, err := client.Get(modifiedLine)
		if err != nil {
			fmt.Println("Error making request:", err)
			continue
		}
		defer resp.Body.Close()
		// Check for redirect status codes
		if resp.StatusCode == http.StatusMovedPermanently || resp.StatusCode == http.StatusFound ||
			resp.StatusCode == http.StatusSeeOther || resp.StatusCode == http.StatusNotModified {
			// Get the Location header
			if location := resp.Header.Get("Location"); location != "" {
				// Print only if the Location header contains the site URL
				if strings.Contains(location, *siteURL) {
					fmt.Printf("Requested URL: %s\nStatus Code: %d\nLocation Header: %s\n", modifiedLine, resp.StatusCode, location)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

// replaceURLParameter replaces the value of the URL parameter in the given line with the new value
func replaceURLParameter(line, newValue string) string {
	parsedURL, err := url.Parse(line)
	if err != nil {
		return line
	}
	queryParams := parsedURL.Query()
	for param := range queryParams {
		queryParams.Set(param, newValue)
	}
	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String()
}
