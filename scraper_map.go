package main

var ScraperMap = map[string]func(string, *Source) ([]Post, error){
	"GoDevScraper":  GoDevScraper,
	"GopherAcademy": nil,
	"FeedScraper":   FeedScraper,
}
