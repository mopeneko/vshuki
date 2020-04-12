package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mopeneko/vshuki/api/database/table"
	"log"
	"math"
	"os"
	"time"
)

var delayCount = 1

const maxCount = 8

func Init() (*gorm.DB, error) {
	db, err := connect()

	if err != nil {
		return nil, err
	}

	migrate(db)

	return db, nil
}

func connect() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_SSL_MODE"),
		),
	)

	if err != nil {
		if delayCount > maxCount {
			return nil, err
		}

		delay := calcDelay(delayCount)
		sleepTime := time.Second * time.Duration(delay)

		log.Printf("failed to initialize database: %+v\nretry after %d seconds...\n", err, delay)
		time.Sleep(sleepTime)

		delayCount++
		return connect()
	}

	return db, nil
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(table.UserAuth{})
	db.AutoMigrate(table.User{})
	db.AutoMigrate(table.Channel{})
	db.AutoMigrate(table.Video{})
	db.AutoMigrate(table.Post{})
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func calcDelay(x int) int {
	return pow(2, x-1)
}
