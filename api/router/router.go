package router

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mopeneko/vshuki/api/controller"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Init(db *gorm.DB) (*echo.Echo, error) {
	e := echo.New()

	jwtSecret, err := getJWTSecret()

	if err != nil {
		return nil, err
	}

	jwtMiddleware := middleware.JWT(jwtSecret)

	baseController := controller.NewBaseController(db)

	postsController := controller.NewPostsController(baseController)
	e.GET("/posts", postsController.GetPosts)

	return e, nil
}

var secretSource = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const secretSize = 128

func getJWTSecret() ([]byte, error) {
	filename := "data/jwt_secret.txt"

	if !isFileExist(filename) {
		file, err := os.Create(filename)

		if err != nil {
			return []byte{}, nil
		}

		defer file.Close()

		secret := generateJWTSecret(secretSize)
		_, err = file.Write(secret)

		if err != nil {
			return []byte{}, nil
		}

		return secret, nil
	}

	secret, err := ioutil.ReadFile(filename)

	if err != nil {
		return []byte{}, err
	}

	return secret, nil
}

func isFileExist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func generateJWTSecret(size uint) []byte {
	secret := make([]byte, size)

	for i := range secret {
		secret[i] = secretSource[rand.Intn(len(secretSource))]
	}

	return secret
}
