package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/mopeneko/vshuki/api/jwt"
	"github.com/mopeneko/vshuki/api/model"
	"github.com/mopeneko/vshuki/api/view"
	"net/http"
)

type AuthController struct {
	*BaseController
	jwtSecret []byte
}

func NewAuthController(base *BaseController, jwtSecret []byte) *AuthController {
	return &AuthController{base, jwtSecret}
}

type PostSignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *AuthController) PostSignIn(ctx echo.Context) error {
	result := new(view.PostSignInResult)

	req := new(PostSignInRequest)
	if err := ctx.Bind(&req); err != nil {
		result.Message = "不正なデータです。"
		return view.RenderPostSignInResult(ctx, http.StatusBadRequest, result)
	}

	if err := validate.Struct(req); err != nil {
		result.Message = "データが不足しています。"
		return view.RenderPostSignInResult(ctx, http.StatusBadRequest, result)
	}

	user := model.AuthModel{}.FindUser(c.db, req.Email)
	if len(user.UUID) <= 0 {
		result.Message = "メールアドレスまたはパスワードが間違っています。"
		return view.RenderPostSignInResult(ctx, http.StatusBadRequest, result)
	}

	err := model.AuthModel{}.ComparePassword(user.UserAuth.Password, []byte(req.Password))
	if err != nil {
		result.Message = "メールアドレスまたはパスワードが間違っています。"
		return view.RenderPostSignInResult(ctx, http.StatusBadRequest, result)
	}

	token, err := jwt.GenerateJWT(user.UUID, c.jwtSecret)
	if err != nil {
		result.Message = "内部エラーが発生しました。"
		return view.RenderPostSignInResult(ctx, http.StatusInternalServerError, result)
	}

	result.Token = token
	return view.RenderPostSignInResult(ctx, http.StatusOK, result)
}

type PostSignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *AuthController) PostSignUp(ctx echo.Context) error {
	result := new(view.PostSignUpResult)

	req := new(PostSignUpRequest)
	if err := ctx.Bind(&req); err != nil {
		result.Message = "不正なデータです。"
		return view.RenderPostSignUpResult(ctx, http.StatusBadRequest, result)
	}

	if err := validate.Struct(req); err != nil {
		result.Message = "データが不足しています。"
		return view.RenderPostSignUpResult(ctx, http.StatusBadRequest, result)
	}

	user := model.AuthModel{}.FindUser(c.db, req.Email)
	if len(user.UUID) > 0 {
		result.Message = "そのメールアドレスは使用済みです。"
		return view.RenderPostSignUpResult(ctx, http.StatusBadRequest, result)
	}

	hashedPassword, err := model.AuthModel{}.HashPassword([]byte(req.Password))
	if err != nil {
		result.Message = "内部エラーが発生しました。"
		return view.RenderPostSignUpResult(ctx, http.StatusInternalServerError, result)
	}

	userID := model.AuthModel{}.CreateUser(c.db, req.Email, hashedPassword)

	token, err := jwt.GenerateJWT(userID, c.jwtSecret)
	if err != nil {
		result.Message = "内部エラーが発生しました。"
		return view.RenderPostSignUpResult(ctx, http.StatusInternalServerError, result)
	}

	result.Token = token
	return view.RenderPostSignUpResult(ctx, http.StatusOK, result)
}
