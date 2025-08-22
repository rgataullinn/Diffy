package main

import (
	"diffy/internal/fetcher"
	"fmt"
	"log"
)

func main() {
	url := "https://www.ulta.com/p/macximal-sleek-satin-lipstick-pimprod2047503?sku=2630257"

	html, err := fetcher.GetHTML(url, "#root")
	if err != nil {
		log.Fatalf("fetch failed: %v", err)
	}

	fmt.Println("Fetched HTML length:", html)
}
