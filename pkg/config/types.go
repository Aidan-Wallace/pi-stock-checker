package config

// TODO: transform yaml names to kebob case

type StockIndicator struct {
	StyleAttributeName string `yaml:"styleAttributeName"`
	Indicator          string `yaml:"indicator"`
}

type Config struct {
	Name                string         `yaml:"name"`
	BaseURL             string         `yaml:"baseUrl"`
	OutOfStockIndicator StockIndicator `yaml:"outOfStockIndicator"`
	InStockIndicator    StockIndicator `yaml:"inStockIndicator"`
}
