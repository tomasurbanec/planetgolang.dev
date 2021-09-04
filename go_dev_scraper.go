package main

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func GoDevScraper(key string, src *Source) ([]Post, error) {
	ary := []Post{}

	body, err := Scrape("https://go.dev/blog/")

	if err != nil {
		return ary, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))

	if err != nil {
		return ary, err
	}

	var lastErr error

	doc.Find(".blogtitle").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")

		if href == "" || href == "/blog/all" {
			return
		}

		var t time.Time
		t, lastErr = time.Parse("_2 January 2006", s.Find(".date").Text())
		url := "https://go.dev" + href
		previousPost, _ := FindPostByUrl(url)

		post := Post{
			Model:       gorm.Model{ID: previousPost.ID},
			Title:       s.Find("a").Text(),
			Author:      s.Find(".author").Text(),
			Url:         url,
			Source:      key,
			PublishedAt: t,
		}

		ary = append(ary, post)
	})

	doc.Find(".blogsummary").Each(func(i int, s *goquery.Selection) {
		ary[i].Summary = strings.TrimSpace(s.Text())
	})

	if lastErr != nil {
		return ary, lastErr
	}

	return ary, nil
}
