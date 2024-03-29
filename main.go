package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"

	"github.com/sebdeveloper6952/blossom-server/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger, err := NewLog(os.Getenv("LOG_LEVEL"))
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
		os.Getenv("CDN_URL"),
		database,
		storage,
		hashing,
	)
	if err != nil {
		log.Fatal(err)
	}

	whitelistedPksSlice := strings.Split(os.Getenv("WHITELISTED_PUBKEYS"), ",")
	whitelistedPks := make(map[string]struct{})
	for i := range whitelistedPksSlice {
		whitelistedPks[whitelistedPksSlice[i]] = struct{}{}
	}

	api := SetupApi(os.Getenv("API_ADDR"), server, whitelistedPks, logger)
	api.Run()
}
