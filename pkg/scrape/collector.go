package scrape

import "github.com/gocolly/colly"

func NewCollector(baseURL string) Collector {
	return Collector{
		Collector: colly.NewCollector(),
		BaseURL:   baseURL,
	}
}
