package table

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID string
	Name   string
	Auth   *UserAuth
}
