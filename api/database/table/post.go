package table

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Video   Video
	VideoID uint
	Comment string
	Poster  *User
}
