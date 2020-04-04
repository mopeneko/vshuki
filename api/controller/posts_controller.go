package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/vshuki/api/model"
	"github.com/mopeneko/vshuki/api/view"
	"net/http"
)

type PostsController struct {
	*BaseController
}

func NewPostsController(base *BaseController) *PostsController {
	return &PostsController{base}
}

func (c *PostsController) GetPosts(ctx echo.Context) error {
	result := &view.GetPostsResult{}

	posts, err := model.PostsModel{}.GetLatestPosts(c.db)

	if err != nil {
		return view.RenderGetPostsResult(ctx, http.StatusInternalServerError, result)
	}

	result.Posts = posts
	return view.RenderGetPostsResult(ctx, http.StatusOK, result)
}

type postPostsRequest struct {
	VideoID string `json:"video_id"`
	Comment string `json:"comment"`
}

func (c *PostsController) PostPosts(ctx echo.Context) error {
	result := new(view.PostPostsResult)

	request := new(postPostsRequest)
	err := ctx.Bind(request)

	if err != nil {
		return view.RenderPostPostsResult(ctx, http.StatusBadRequest, result)
	}

	err = model.PostsModel{}.StorePost(c.db, request.VideoID, request.Comment, nil)

	if err != nil {
		return view.RenderPostPostsResult(ctx, http.StatusInternalServerError, result)
	}

	return view.RenderPostPostsResult(ctx, http.StatusOK, result)
}
