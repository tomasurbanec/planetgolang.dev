package main

import (
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	InitializeDb()

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "scrape",
				Usage:  "scrapes everything",
				Action: ScrapeCommand,
			},
			{
				Name:   "generate",
				Usage:  "generates the site",
				Action: Generate,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
