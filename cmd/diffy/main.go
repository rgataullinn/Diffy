package main

import (
	"diffy/internal/fetcher"
	"diffy/internal/storage"
	"log"
)

func main() {
	url := "https://www.ulta.com/p/macximal-sleek-satin-lipstick-pimprod2047503?sku=2630257"

	html, err := fetcher.GetHTML(url, "#root")
	if err != nil {
		log.Fatalf("fetch failed: %v", err)
	}

	storage.SaveHTML("root", html[0])

	// fmt.Println("Fetched HTML length:", html)
}
