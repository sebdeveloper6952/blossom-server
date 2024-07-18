package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"

	ginApi "github.com/sebdeveloper6952/blossom-server/api/gin"
	"github.com/sebdeveloper6952/blossom-server/db"
	blobDescriptorRepos "github.com/sebdeveloper6952/blossom-server/repos/blob_descriptor"
	"github.com/sebdeveloper6952/blossom-server/services"
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

	blobDescriptorRepo, err := blobDescriptorRepos.NewSqlcRepo(
		database,
		config.CdnUrl,
		logger,
	)
	if err != nil {
		log.Fatal(err)
	}

	hasher, err := services.NewSha256()
	if err != nil {
		log.Fatal(err)
	}

	whitelistedPks := make(map[string]struct{})
	for i := range config.WhitelistedPubkeys {
		whitelistedPks[config.WhitelistedPubkeys[i]] = struct{}{}
	}

	api := ginApi.SetupApi(
		blobDescriptorRepo,
		hasher,
		config.CdnUrl,
		config.ApiAddr,
		whitelistedPks,
		logger,
	)
	api.Run()
}
