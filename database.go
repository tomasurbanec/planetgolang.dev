package planetgolang

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InsertPost(p Post) error {
	db, err := sql.Open("sqlite3", "./planetgolang.db")

	if err != nil {
		return err
	}

	stmt, err := db.Prepare(`
		INSERT INTO posts
		(title, summary, author, source, published_at, url)
		VALUES
		(?, ?, ?, ?, ?, ?)
		ON CONFLICT(url) DO UPDATE
		set title = excluded.title,
		summary = excluded.summary,
		author = excluded.author,
		source = excluded.source,
		published_at = excluded.published_at;
	`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		p.Title,
		p.Summary,
		p.Author,
		p.Source,
		p.PublishedAt,
		p.Url,
	)

	if err != nil {
		return err
	}

	return nil
}
