package main

import (
	"fmt"
	"os"
)

func main() {
	// CLI arguments without program path
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Start crawling
	base_url := args[0]
	fmt.Printf("starting crawl of: %v", base_url)

	// Get HTML test response
	response, err := getHTML(base_url)
	if err != nil {
		fmt.Errorf("Failed to get HTML: %v", err)
	}
	fmt.Printf("Response: %s\n", response)
}