package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Parse both URLs into URL structs
	baseURL, err := url.Parse(rawBaseURL)
	currentURL, err := url.Parse(rawCurrentURL)

	// Ensure currentURL is in the same domain as baseURL
	if baseURL.Host != currentURL.Host {
		// Different domains
		fmt.Errorf("ERROR: currentURL outside of baseURL's domain, %s, %s", baseURL.String(), currentURL.String())
		return
	}

	// Normalize rawCurrentURL
	normalizedCurrentURL, err := normalizeURL(currentURL.String())
	if err != nil {
		fmt.Errorf("ERROR: Couldn't normalize URL: %v", err)
		return
	}

	// If pages map already contains normalizedCurrentURL, just
	// increment its count and return
	// Else, add to map and set count to 1, then continue
	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL] += 1
	} else {
		pages[normalizedCurrentURL] = 1
	}

	// Get HTML of current URL, print it
	html, err := getHTML(normalizedCurrentURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't get HTML of page %s, reason: %v", normalizedCurrentURL, err)
		return
	}
	fmt.Printf("%s", html)

	// Get all of the URLS from the html
	parsedRawCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't parse rawCurrentURL: %v", err)
		return
	}

	urls, err := getURLsFromHTML(html, parsedRawCurrentURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't get URLs from HTML: %v", err)
		return
	}

	// BASE CASE: If there are no URLs, return
	if len(urls) <= 0 {
		return
	}

	// Recursively call crawlPage on all URLs found in that html
	for _, url := range urls {
		crawlPage(baseURL.String(), url, pages)
	}
}