package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	pl "github.com/paweljw/planetgolang"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:   "initializedb",
				Usage:  "Primes local SQLite instance",
				Action: initializeDb,
			},
			{
				Name:   "scrape",
				Usage:  "scrapes everything",
				Action: scrape,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initializeDb(_ *cli.Context) error {
	db, err := sql.Open("sqlite3", "./planetgolang.db")

	if err != nil {
		return err
	}

	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS posts (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				title TEXT,
				summary TEXT,
				author VARCHAR(255),
				source VARCHAR(255),
				published_at DATETIME,
				url VARCHAR(255) UNIQUE
			)
		`)

	if err != nil {
		return err
	}

	return nil
}

func scrape(_ *cli.Context) error {
	ary, err := pl.GoDevScraper()

	if err != nil {
		return err
	}

	for _, post := range ary {
		fmt.Printf("%s by %s / %s @ %s on %s\n", post.Title, post.Author, post.Source, post.Url, post.PublishedAt)
		fmt.Println(post.Summary)

		err := pl.InsertPost(post)

		if err != nil {
			return err
		}
	}

	return nil
}
