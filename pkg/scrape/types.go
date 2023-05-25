package scrape

import "github.com/gocolly/colly"

type ScrapeResult struct {
	Result *colly.HTMLElement
	SomethingWentWrong bool
}

type Collector struct {
	Collector *colly.Collector
	BaseURL   string
}
