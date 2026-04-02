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
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}

	// List of all links
	links := []string{}

	// Each() runs a given fuction whenever it matches the Find() param
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		link, _ := s.Attr("href")

		// Relative paths to absolute
		if strings.HasPrefix(link, "/") {
			link = baseURL.String() + link
		}

		links = append(links, link)
	})

	return links, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return []string{}, err
	}

	images := []string{}

	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		image, _ := s.Attr("src")

		if strings.HasPrefix(image, "/") {
			image = baseURL.String() + image
		}

		images = append(images, image)
	})

	return images, nil
}