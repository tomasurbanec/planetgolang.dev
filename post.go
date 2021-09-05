package main

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string
	Summary     string
	Url         string
	Author      string
	Source      string
	PublishedAt time.Time
	SourceUrl   string
}

func (p *Post) FormattedShortPublishedAt() string {
	return p.PublishedAt.Format("02 Jan 06")
}

func (p *Post) FormattedPublishedAt() string {
	return p.PublishedAt.Format("02 Jan 06 15:04 MST")
}
