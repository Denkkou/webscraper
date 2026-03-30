package main

import(
	"github.com/PuerkitoBio/goquery"
	"strings"
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