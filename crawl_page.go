package main

import (

)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	// Ensure rawCurrentURL is in the same domain as rawBaseURL
	// else return

	// Normalize rawCurrentURL
	//normalizedURL := normalizeURL(rawCurrentURL)

	// If pages map already contains normalizedURL, just
	// increment its count and return
	// Else, add to map and set count to 1, then continue

	// Get HTML of current URL, print it

	// Recursively call crawlPage on all URLs found in that html
}