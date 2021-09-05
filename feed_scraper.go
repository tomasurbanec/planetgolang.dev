package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
	"gorm.io/gorm"
)

func FeedScraper(key string, src *Source) ([]Post, error) {
	fp := gofeed.NewParser()

	ary := []Post{}

	feedBody, err := Scrape(src.ScrapeUrl)

	if err != nil {
		return ary, err
	}

	feed, err := fp.ParseString(feedBody)

	if err != nil {
		return ary, err
	}

	for _, item := range feed.Items {
		authors := []string{}

		for _, author := range item.Authors {
			authors = append(authors, author.Name)
		}

		previousPost, _ := FindPostByUrl(item.Link)

		if !previousPost.DeletedAt.Time.IsZero() {
			continue
		}

		post := Post{
			Model:       gorm.Model{ID: previousPost.ID},
			Title:       item.Title,
			Summary:     item.Description,
			Url:         item.Link,
			Author:      strings.Join(authors, ", "),
			Source:      key,
			PublishedAt: *item.PublishedParsed,
		}

		ary = append(ary, post)
	}

	return ary, nil
}
