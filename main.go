package main

import (
	"log"
	"os"
	"strconv"

	cli "github.com/urfave/cli/v2"
	"gorm.io/gorm"
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
			{
				Name:  "delete",
				Usage: "deletes post by id",
				Action: func(c *cli.Context) error {
					sid := c.Args().Get(0)

					if sid == "" {
						log.Fatal("Can't delete empty id")
					}

					id, _ := strconv.ParseUint(sid, 10, 64)

					post := Post{Model: gorm.Model{ID: uint(id)}}

					err := Db.Find(&post).Error

					if err != nil {
						return err
					}

					err = Db.Delete(&post).Error

					return err
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
