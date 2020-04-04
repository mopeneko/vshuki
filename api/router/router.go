package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init(db *gorm.DB) *echo.Echo {
	return echo.New()
}
