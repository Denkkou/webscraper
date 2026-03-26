package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	fullPath := parsedURL.Host + parsedURL.Path
	normalizedURL := strings.ToLower(strings.TrimSuffix(fullPath, "/"))
	
	return normalizedURL, nil
}