package main

import(
	//"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	// Scan string for <h1></h1> to get the title
	// If there is no <h1>, scan for <h2> instead
	// Else return empty string (no heading)
	return ""
}

func getFirstParagraphFromHTML(html string) string {
	// Scan for <main> tag and then check if <p> within
	// If so, return the contents of that <p>
	// Else, go back to the top and find the first instance of <p>
	// Return empty if no <p> is found
	return ""
}