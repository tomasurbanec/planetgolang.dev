package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

func ScrapeCommand(_ *cli.Context) error {
	sources := make(map[string]Source)

	feeds, err := os.ReadFile("./feeds.yml")

	if err != nil {
		return err
	}

	yaml.Unmarshal(feeds, &sources)

	for key, source := range sources {
		log.Printf("Processing source %s", key)

		scraper := ScraperMap[source.Scraper]

		if scraper != nil {
			posts, err := scraper(key, &source)

			if err != nil {
				log.Printf("%s failed: %s", key, err.Error())
			} else {
				for _, post := range posts {
					err := InsertPost(post)

					if err != nil {
						log.Printf("Insert failed: %s", err.Error())
					}
				}
			}
		} else {
			log.Printf("%s scraper missing: %s", key, source.Scraper)
		}
	}

	return nil
}
