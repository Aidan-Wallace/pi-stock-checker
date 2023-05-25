package main

import (
	"flag"
	"log"
	"pi-stock-checker/pkg/config"
	"pi-stock-checker/pkg/scrape"
	"time"
)

const CONFIG_DIR = "./config"

func loadConfigFiles(dp string) []config.Config {
	configs := config.LoadConfigFiles(dp)
	if len(configs) == 0 {
		log.Fatal("No config files found")
	}
	return configs
}

type SomethingWentWrong struct {
	Message    string
	DidGoWrong bool
}

type Result struct {
	DateFinished       time.Time
	Name               string               `json:"name"`
	Url                string               `json:"url"`
	IsInStock          bool                 `json:"isInStock"`
	IsOutOfStock       bool                 `json:"isOutOfStock"`
	SomethingWentWrong []SomethingWentWrong `json:"somethingWentWrong"`
}

func doScrape(baseURL string, i string) scrape.ScrapeResult {
	scraper := scrape.NewCollector(baseURL)
	result := scrape.ScrapeResult{}

	scrape.Run(&scraper, i, &result)

	return result
}

func getStock(r scrape.ScrapeResult, t string) bool {
	if r.Result == nil {
		return false
	}

	return r.Result.Text == t
}

func run(fp string) []Result {
	results := []Result{}

	configs := loadConfigFiles(fp)
	for _, conf := range configs {
		outOfStock := doScrape(conf.BaseURL, conf.OutOfStockIndicator.StyleAttributeName)
		inStock := doScrape(conf.BaseURL, conf.InStockIndicator.StyleAttributeName)

		results = append(results, Result{
			DateFinished: time.Now(),
			Name:         conf.Name,
			Url:          conf.BaseURL,
			IsInStock:    getStock(inStock, conf.InStockIndicator.Indicator),
			IsOutOfStock: getStock(outOfStock, conf.OutOfStockIndicator.Indicator),
			SomethingWentWrong: []SomethingWentWrong{
				{
					Message:    "Out of stock",
					DidGoWrong: outOfStock.SomethingWentWrong,
				},
				{
					Message:    "In stock",
					DidGoWrong: inStock.SomethingWentWrong,
				},
			},
		})
	}

	return results
}

var configsDirPath = flag.String("cd", CONFIG_DIR, "Dir which holds yaml config files")
var outputFilePath = flag.String("of", "results.json", "Json file to export results to")

func main() {
	flag.Parse()

	results := run(*configsDirPath)
	config.ExportJSON(*outputFilePath, results)
}
