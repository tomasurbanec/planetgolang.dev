package planetgolang

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GoDevScraper() ([]Post, error) {
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

		post := Post{
			Title:       s.Find("a").Text(),
			Author:      s.Find(".author").Text(),
			Url:         "https://go.dev" + href,
			Source:      "The Go Blog",
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
