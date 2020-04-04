package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mopeneko/vshuki/api/database/table"
)

type PostsModel struct{}

func (PostsModel) GetLatestPosts(db *gorm.DB) ([]*table.Post, error) {
	return []*table.Post{}, nil
}

func (PostsModel) StorePost(db *gorm.DB, videoID string, comment string, user *table.User) error {
	post := &table.Post{
		VideoID: videoID,
		Comment: comment,
		Poster: user,
	}

	db.Create(post)
	return nil
}
