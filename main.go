package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sebdeveloper6952/blossom-server/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.NewDB(
		os.Getenv("DB_PATH"),
		os.Getenv("DB_MIGRATIONS_PATH"),
	)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := NewFsStorage("media")
	if err != nil {
		log.Fatal(err)
	}

	hashing, err := NewSha256()
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(
		database,
		storage,
		hashing,
	)
	if err != nil {
		log.Fatal(err)
	}

	api := SetupApi("127.0.0.1:8000", server)
	api.Run()
}
