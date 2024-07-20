package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	ginApi "github.com/sebdeveloper6952/blossom-server/api/gin"
	"github.com/sebdeveloper6952/blossom-server/db"
	blobDescriptorRepos "github.com/sebdeveloper6952/blossom-server/repos/blob_descriptor"
)

func main() {
	config, err := NewConfig("config.yml")
	if err != nil {
		fmt.Printf("load config: %v", err)
		os.Exit(1)
	}

	logger, err := NewLog(config.LogLevel)
	if err != nil {
		log.Printf("create logger: %v", err)
		os.Exit(1)
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

	whitelistedPks := make(map[string]struct{})
	for i := range config.WhitelistedPubkeys {
		whitelistedPks[config.WhitelistedPubkeys[i]] = struct{}{}
	}

	api := ginApi.SetupApi(
		blobDescriptorRepo,
		config.CdnUrl,
		config.ApiAddr,
		whitelistedPks,
		logger,
	)
	api.Run()
}
