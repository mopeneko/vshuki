package model

import (
	"crypto/sha512"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/mopeneko/vshuki/api/database/table"
	"golang.org/x/crypto/bcrypt"
)

type AuthModel struct{}

func (AuthModel) HashPassword(password []byte) (string, error) {
	hash := sha512.Sum512(password)
	encryptedHash, err := bcrypt.GenerateFromPassword(hash[:], 10)

	if err != nil {
		return "", err
	}

	return string(encryptedHash), nil
}

func (AuthModel) ComparePassword(hash string, password []byte) error {
	passwordHash := sha512.Sum512(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), passwordHash[:])
	return err
}

func (AuthModel) CreateUser(db *gorm.DB, email, password string) string {
	user := &table.User{
		UUID: uuid.New().String(),
		Name: "名無しさん",
		UserAuth: table.UserAuth{
			Email:    email,
			Password: password,
		},
	}

	db.Create(user)

	return user.UUID
}

func (AuthModel) FindUser(db *gorm.DB, email string) *table.User {
	user := new(table.User)

	db.Preload("UserAuth", "email = ?", email).First(user)

	return user
}
