package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var Db *gorm.DB

const PER_PAGE = 10

func InitializeDb() error {
	var err error
	Db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	Db.AutoMigrate(&Post{})

	return err
}

func InsertPost(p Post) error {
	err := Db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&p).Error

	if err != nil {
		return err
	}

	return nil
}

func ReadPosts(page int) ([]Post, error) {
	posts := []Post{}

	res := Db.Limit(PER_PAGE).Offset(page * PER_PAGE).Order("published_at desc").Find(&posts)

	return posts, res.Error
}

func CountPosts() (int64, error) {
	posts := []Post{}

	res := Db.Find(&posts)

	return res.RowsAffected, res.Error
}
