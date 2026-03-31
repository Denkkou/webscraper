package main

import(
	"github.com/PuerkitoBio/goquery"
	"strings"
	"net/url"
)

func getHeadingFromHTML(html string) (string, error) {
	// Create document reader to parse input
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	// Find header tag in document
	header := doc.Find("h1").Text()
	if header == "" {
		// <h2> as a fallback
		header = doc.Find("h2").Text()
	}

	// If still empty, return anyway
	return header, nil
}

func getFirstParagraphFromHTML(html string) (string, error) {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}

	paragraph := ""
	main := doc.Find("main")

	// If <main> tag is present
	if main.Length() > 0 {
		// Find first <p> tag within it
		paragraph = main.Find("p").First().Text()
	}

	// If paragraph is empty, default to first <p> tag in doc
	if paragraph == "" {
		paragraph = doc.Find("p").First().Text()
	}

	// If still empty, return anyway
	return paragraph, nil
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	// Parse the htmlBody for instances of URLs and also
	// for relative paths, such as /logo.png etc.

	// Create a list of all of these instances
	// including all relative paths being turned into
	// absolute ones, such as www.crawler.com/logo.png

	// Make sure to find every <a> tag

	return []string{""}, nil
}