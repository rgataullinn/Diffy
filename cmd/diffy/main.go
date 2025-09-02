package main

import (
	"diffy/internal/fetcher"
	"diffy/internal/storage"
	"log"
)

type url_data struct {
	id  string
	url string
}

func main() {
	data := map[string][]url_data{
		"ulta": {
			{"pimprod2047503", "https://www.ulta.com/p/macximal-sleek-satin-lipstick-pimprod2047503?sku=2630257"},
			{"xlsImpprod10792007", "https://www.ulta.com/p/almost-lipstick-xlsImpprod10792007?sku=2153395"},
			{"pimprod2044074", "https://www.ulta.com/p/dream-coat-curly-hair-pimprod2044074?sku=2647829"},
		},
	}

	for client, urls := range data {
		for _, url_data := range urls {
			html, err := fetcher.GetHTML(url_data.url, "#root")
			if err != nil {
				log.Fatalf("fetch failed: %v", err)
			}

			storage.SaveHTML("root", html[0], "snapshots/"+client+"/"+url_data.id)
		}
	}
}
