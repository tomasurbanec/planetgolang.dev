package planetgolang

import (
	"time"
)

type Post struct {
	Id          int64
	Title       string
	Summary     string
	Url         string
	Author      string
	Source      string
	PublishedAt time.Time
}
