package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Parse both URLs into URL structs
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't parse baseURL: %v", err)
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't parse currentURL: %v", err)
		return
	}

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
		fmt.Println("Page already present in pages map", normalizedCurrentURL)
		return
	}
	
	pages[normalizedCurrentURL] = 1

	// Get HTML of current URL, print it
	html, err := getHTML(currentURL.String())
	if err != nil {
		fmt.Errorf("ERROR: Couldn't get HTML of page %s, reason: %v", currentURL.String(), err)
		return
	}
	//fmt.Printf("\nPage HTML:\n%s", html)

	// Get all of the URLS from the html
	urls, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Errorf("ERROR: Couldn't get URLs from HTML: %v", err)
		return
	}

	// If there are no URLs, return
	if len(urls) <= 0 {
		fmt.Println("No URLs found in HTML belonging to ", currentURL)
		return
	}

	// Recursively call crawlPage on all URLs found in that html
	for _, url := range urls {
		crawlPage(baseURL.String(), url, pages)
	}
}