package controller

import "github.com/jinzhu/gorm"

type BaseController struct {
	db *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{db}
}
