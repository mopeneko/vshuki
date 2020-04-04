package model

import (
	"github.com/jinzhu/gorm"
	"github.com/mopeneko/vshuki/api/database/table"
)

type PostsModel struct{}

func (PostsModel) GetLatestPosts(db *gorm.DB) ([]*table.Post, error) {
	return []*table.Post{}, nil
}
