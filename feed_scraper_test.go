package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const rssFeed = `<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0">
  <channel>
    <title>Go Programming Language News</title>
    <link>https://example.com</link>
    <description>Latest updates and articles about the Go programming language</description>
    <language>en-us</language>
    <lastBuildDate>Tue, 27 Feb 2025 12:00:00 GMT</lastBuildDate>

    <item>
      <title>Go 1.22 Released: What’s New?</title>
      <link>https://example.com/go-1-22</link>
      <description>The latest version of Go brings improved error handling, enhanced generics, and better performance.</description>
      <pubDate>Tue, 27 Feb 2025 11:00:00 GMT</pubDate>
      <guid>https://example.com/go-1-22</guid>
    </item>

    <item>
      <title>Concurrency in Go: Best Practices for Goroutines</title>
      <link>https://example.com/go-concurrency</link>
      <description>Learn how to effectively use goroutines, channels, and context for building concurrent applications in Go.</description>
      <pubDate>Tue, 27 Feb 2025 10:00:00 GMT</pubDate>
      <guid>https://example.com/go-concurrency</guid>
    </item>

    <item>
      <title>Building REST APIs with Go and Gin</title>
      <link>https://example.com/go-gin-rest-api</link>
      <description>A step-by-step guide to building fast and lightweight REST APIs using the Gin framework in Go.</description>
      <pubDate>Tue, 27 Feb 2025 09:00:00 GMT</pubDate>
      <guid>https://example.com/go-gin-rest-api</guid>
    </item>
  </channel>
</rss>`

func TestFeedScraper(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(rssFeed))
	}))
	defer server.Close()

	src := Source{
		Key:       "foo",
		Title:     "Sample Go blog title",
		ScrapeUrl: server.URL,
		Url:       server.URL,
		Scraper:   "FeedScraper",
	}

	posts, _ := FeedScraper("foo", &src)

	if len(posts) != 3 {
		t.Error("FeedScraper didn't get all the posts")
	}

	urls := [3]string{
		"https://example.com/go-1-22",
		"https://example.com/go-concurrency",
		"https://example.com/go-gin-rest-api",
	}

	for index, url := range urls {
		post := posts[index]

		if post.Url != url {
			t.Errorf("FeedScraper expected %s, but got %s", post.Url, url)
		}

		if post.Source != src.Key {
			t.Errorf("FeedScraper is returning the wrong source, expected %s, but got %s",
				post.Source,
				src.Key)
		}
	}
}
