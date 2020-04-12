package view

import "github.com/labstack/echo/v4"

type PostSignInResult struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

func RenderPostSignInResult(ctx echo.Context, code int, result *PostSignInResult) error {
	return ctx.JSON(code, result)
}

type PostSignUpResult struct {
	Token string `json:"token"`
	Message string `json:"message"`
}

func RenderPostSignUpResult(ctx echo.Context, code int, result *PostSignUpResult) error {
	return ctx.JSON(code, result)
}
