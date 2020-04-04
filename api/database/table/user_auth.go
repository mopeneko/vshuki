package table

import "github.com/jinzhu/gorm"

type UserAuth struct {
	gorm.Model
	Email    string
	Password string
}
