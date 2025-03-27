package main

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	dateUnderTest := "2022-Feb-02"
	publishedAt, _ := time.Parse("2006-Jan-02", dateUnderTest)
	post := &Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	}

	expectedFormattedShortPublishedAt := "02 Feb 22"
	receivedFormattedShortPublishedAt := post.FormattedShortPublishedAt()

	if receivedFormattedShortPublishedAt != expectedFormattedShortPublishedAt {
		t.Errorf("Post's short published at date is incorrect, expected: %s, received: %s.",
			expectedFormattedShortPublishedAt,
			receivedFormattedShortPublishedAt)
	}

	expectedFormattedPublishedAt := "02 Feb 22 00:00 UTC"
	receivedFormattedPublishedAt := post.FormattedPublishedAt()

	if receivedFormattedPublishedAt != expectedFormattedPublishedAt {
		t.Errorf("Post's formatted published at date is incorrect, expected: %s, received: %s.",
			expectedFormattedPublishedAt,
			receivedFormattedPublishedAt)
	}
}
