package main

import (
	"html/template"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	fd "github.com/gorilla/feeds"

	strip "github.com/grokify/html-strip-tags-go"
	cli "github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

type Page struct {
	Posts       []Post
	TotalPages  int
	CurrentPage int
}

func Generate(_ *cli.Context) error {
	indexTmpl := template.New("index.html.tmpl")

	indexTmpl.Funcs(template.FuncMap{
		"minus": func(a, b int) int {
			return a - b
		},
		"plus": func(a, b int) int {
			return a + b
		},
	})

	indexTmpl.ParseGlob("./templates/*.tmpl")

	currentPage := 0

	totalPosts, err := CountPosts()

	if err != nil {
		return err
	}

	totalPages := int(math.Ceil(float64(totalPosts) / float64(PER_PAGE)))

	sources := make(map[string]Source)

	feeds, err := os.ReadFile("./feeds.yml")

	if err != nil {
		return err
	}

	f, err := os.Create("./dist/what.html")
	if err != nil {
		log.Fatalf("Failed to generate site: %s", err.Error())
	}
	defer f.Close()

	indexTmpl.ExecuteTemplate(f, "what.tmpl", nil)

	yaml.Unmarshal(feeds, &sources)

	now := time.Now()
	feed := &fd.Feed{
		Title:       "Planet Golang",
		Link:        &fd.Link{Href: "https://planetgolang.dev"},
		Description: "An unopinionated collection of newest Golang articles from all around the web.",
		Author:      &fd.Author{Name: "Paweł J. Wal", Email: "hello@planetgolang.dev"},
		Created:     now,
	}

	feedItems := []*fd.Item{}

	for {
		posts, err := ReadPosts(currentPage)

		for i := range posts {
			posts[i].Summary = strip.StripTags(posts[i].Summary)
			if len(posts[i].Summary) > 280 {
				posts[i].Summary = strings.TrimSpace(posts[i].Summary[0:280]) + "..."
			}

			for key, source := range sources {
				if posts[i].Source == key {
					posts[i].Source = source.Title
					posts[i].SourceUrl = source.Url
					break
				}
			}

			feedItems = append(feedItems, &fd.Item{
				Title:       posts[i].Title,
				Link:        &fd.Link{Href: posts[i].Url},
				Description: posts[i].Summary,
				Author:      &fd.Author{Name: posts[i].Author},
				Created:     posts[i].PublishedAt,
			})
		}

		page := Page{Posts: posts, TotalPages: totalPages, CurrentPage: currentPage + 1}

		if err != nil {
			log.Fatalf("Failed to generate site: %s", err.Error())
		}

		if len(posts) == 0 {
			break
		}

		if currentPage == 0 {
			f, err := os.Create("./dist/index.html")
			if err != nil {
				log.Fatalf("Failed to generate site: %s", err.Error())
			}
			defer f.Close()

			err = indexTmpl.ExecuteTemplate(f, "index.html.tmpl", page)

			if err != nil {
				log.Fatal(err.Error())
			}

			f, err = os.Create("./dist/1.html")

			if err != nil {
				log.Fatalf("Failed to generate site: %s", err.Error())
			}
			defer f.Close()

			err = indexTmpl.ExecuteTemplate(f, "index.html.tmpl", page)

			if err != nil {
				log.Fatal(err.Error())
			}

			feed.Items = feedItems
			atom, err := feed.ToAtom()
			if err != nil {
				log.Fatal(err)
			}
			f, err = os.Create("./dist/index.xml")
			if err != nil {
				log.Fatalf("Failed to generate site: %s", err.Error())
			}
			defer f.Close()

			f.Write([]byte(atom))
		} else {
			cPage := currentPage + 1
			f, err := os.Create("./dist/" + strconv.Itoa(cPage) + ".html")

			if err != nil {
				log.Fatalf("Failed to generate site: %s", err.Error())
			}
			defer f.Close()

			err = indexTmpl.ExecuteTemplate(f, "index.html.tmpl", page)

			if err != nil {
				log.Fatal(err.Error())
			}
		}

		currentPage += 1
	}

	return nil
}
