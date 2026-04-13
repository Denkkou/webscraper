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

	// Recursive call
	pages := make(map[string]int)
	crawlPage(base_url, base_url, pages)

	// Print keys and values of the pages map
	for key, val := range pages {
		fmt.Printf("\n%s, %d", key, val)
	}
}