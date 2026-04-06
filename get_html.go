package main 

import (
	"net/http"
	"fmt"
	"io"
)

func getHTML(rawURL string) (string, error) {
	client := http.Client{}
	
	// Build request, set header
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// Catch error codes
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("ERROR: Response status code %v", res.StatusCode) 
	}

	// Catch non-text/html content
	

	// Catch others


	// Read data
	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return "", nil
}