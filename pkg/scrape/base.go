package scrape

import (
	"github.com/gocolly/colly"
)

// c: scrape.Collector [scraping framework config], t: string [html style target], r: ScrapeResult [result of scrape]
func Run(c *Collector, t string, r *ScrapeResult) {
	c.Collector.OnHTML(t, func(e *colly.HTMLElement) {
		r.Result = e
	})

	c.Collector.Visit(c.BaseURL)

	r.SomethingWentWrong = r.Result == nil
}
