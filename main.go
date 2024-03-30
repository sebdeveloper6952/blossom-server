package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/sebdeveloper6952/blossom-server/db"
	"log"
)

func main() {
	config, err := NewConfig("config.yml")

	logger, err := NewLog(config.LogLevel)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := db.NewDB(
		config.Db.Path,
		config.Db.MigrationDir,
	)
	if err != nil {
		log.Fatal(err)
	}

	storage, err := NewFsStorage(config.Storage.BasePath)
	if err != nil {
		log.Fatal(err)
	}

	hashing, err := NewSha256()
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(
		config.CdnUrl,
		database,
		storage,
		hashing,
	)
	if err != nil {
		log.Fatal(err)
	}

	whitelistedPks := make(map[string]struct{})
	for i := range config.WhitelistedPubkeys {
		whitelistedPks[config.WhitelistedPubkeys[i]] = struct{}{}
	}

	api := SetupApi(
		config.ApiAddr,
		server,
		whitelistedPks,
		logger,
	)
	api.Run()
}
