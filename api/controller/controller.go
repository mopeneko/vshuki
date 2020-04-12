package controller

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type BaseController struct {
	db *gorm.DB
}

func NewBaseController(db *gorm.DB) *BaseController {
	return &BaseController{db}
}
