package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/vshuki/api/controller"
)

func Init(db *gorm.DB) *echo.Echo {
	e := echo.New()

	baseController := controller.NewBaseController(db)

	postsController := controller.NewPostsController(baseController)
	e.GET("/posts", postsController.GetPosts)

	return e
}
