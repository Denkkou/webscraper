package main

import (
	"net/url"
	"log"
)

type PageData struct {
	URL				string
	Heading			string
	FirstParagraph	string
	OutgoingLinks	[]string
	ImageURLs		[]string
}

func extractPageData(html, pageURL string) PageData {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		log.Printf("NON-FATAL: Error with parsing pageURL: %v", err)
		return PageData{}
	}

	heading, err := getHeadingFromHTML(html)
	if err != nil {
		log.Printf("NON-FATAL: Error getting heading: %v", err)
	}

	firstp, err := getFirstParagraphFromHTML(html)
	if err != nil {
		log.Printf("NON-FATAL: Error getting first paragraph: %v", err)
	}

	links, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		log.Printf("NON-FATAL: Error getting URLs: %v", err)
	}

	images, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		log.Printf("NON-FATAL: Error getting images: %v", err)
	}

	return PageData{
		URL: 			pageURL,
		Heading: 		heading,
		FirstParagraph: firstp,
		OutgoingLinks: 	links,
		ImageURLs: 		images,
	}
}