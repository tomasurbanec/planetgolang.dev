package main

type Source struct {
	Key       string
	Title     string
	ScrapeUrl string `yaml:"scrape_url"`
	Url       string
	Scraper   string
}
