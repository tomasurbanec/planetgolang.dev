package main

import (
	"errors"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const dbFile = "test.db"

func setupTest(t *testing.T) func(t *testing.T) {
	err := InitializeDb(dbFile, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		t.Error("Error initialising the database")
	}

	return func(t *testing.T) {
		if _, err := os.Stat(dbFile); err == nil {
			dbRemovalError := os.Remove(dbFile)

			if dbRemovalError != nil {
				t.Error("Error removing the database", dbRemovalError)
			}
		}
	}
}

func TestInitializeDb(t *testing.T) {
	teardownSuite := setupTest(t)
	defer teardownSuite(t)

	if _, err := os.Stat(dbFile); errors.Is(err, os.ErrNotExist) {
		t.Error("Error didn't create the database file")
	}
}

func TestFindPostByUrl(t *testing.T) {
	teardownSuite := setupTest(t)
	defer teardownSuite(t)

	publishedAt, _ := time.Parse("2006-Jan-02", "2022-Feb-02")
	Db.Create(&Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	})
	post, _ := FindPostByUrl("http://foo.bar")

	if post.ID != 1 {
		t.Error("Couldn't find the post in the database")
	}
}

func TestInsertPost(t *testing.T) {
	teardownSuite := setupTest(t)
	defer teardownSuite(t)

	_, err := FindPostByUrl("http://foo.bar")

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("Post found when none were inserted")
	}

	publishedAt, _ := time.Parse("2006-Jan-02", "2022-Feb-02")
	post := Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	}
	InsertPost(post)
	insertedPost, _ := FindPostByUrl("http://foo.bar")

	if insertedPost.ID != 1 {
		t.Error("Couldn't find the post in the database")
	}
}

func TestReadPosts(t *testing.T) {
	teardownSuite := setupTest(t)
	defer teardownSuite(t)

	publishedAt, _ := time.Parse("2006-Jan-02", "2022-Feb-02")
	post := Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	}
	InsertPost(post)

	posts, _ := ReadPosts(1)

	if len(posts) != 0 {
		t.Error("Too many posts returned")
	}
}

func TestCountPosts(t *testing.T) {
	teardownSuite := setupTest(t)
	defer teardownSuite(t)

	publishedAt, _ := time.Parse("2006-Jan-02", "2022-Feb-02")
	post := Post{
		Title:       "title",
		Summary:     "bar",
		Url:         "http://foo.bar",
		Author:      "Foo Bar",
		Source:      "source",
		PublishedAt: publishedAt,
		SourceUrl:   "http://foo.bar/post/1",
	}
	InsertPost(post)
	count, _ := CountPosts()

	if count != 1 {
		t.Error("Post count is incorrect")
	}
}
