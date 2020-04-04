package view

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/vshuki/api/database/table"
)

type GetPostsResult struct {
	Posts []*table.Post `json:"posts"`
}

func RenderGetPostsResult(ctx echo.Context, code int, result *GetPostsResult) error {
	return ctx.JSON(code, result)
}
