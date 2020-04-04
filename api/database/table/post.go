package table

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	VideoID string
	Comment string
	Poster  *User
}
