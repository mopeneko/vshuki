package table

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UUID       string
	Name       string
	UserAuth   UserAuth
	UserAuthID uint
}
