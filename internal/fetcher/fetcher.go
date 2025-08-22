package fetcher

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func GetHTML(url string, selectors ...string) ([]string, error) {
	results := []string{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Scraping:", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if len(selectors) == 0 {
		c.OnResponse(func(r *colly.Response) {
			results = append(results, string(r.Body))
		})
	} else {
		for _, sel := range selectors {
			s := sel
			c.OnHTML(s, func(e *colly.HTMLElement) {
				html, err := e.DOM.Html()
				if err == nil {
					results = append(results, html)
				}
			})
		}
	}

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return results, nil
}
