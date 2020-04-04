package main

import (
	"github.com/joho/godotenv"
	"github.com/mopeneko/vshuki/api/database"
	"github.com/mopeneko/vshuki/api/router"
	"log"
)

func main() {
	err := loadDotEnv()

	if err != nil {
		log.Fatalf("Failed to load .env file: %+v\n", err)
	}

	db, err := database.Init()

	if err != nil {
		log.Fatalf("Failed to connect to database: %+v\n", err)
	}

	e := router.Init(db)

	e.Logger.Fatal(e.Start(":4000"))
}

func loadDotEnv() error {
	return godotenv.Load()
}
