package main

import (
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	ginApi "github.com/sebdeveloper6952/blossom-server/api/gin"
	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/logging"
	"github.com/sebdeveloper6952/blossom-server/storage"
)

func main() {
	conf, err := config.NewConfig("config.yml")
	if err != nil {
		log.Printf("read config: %v", err)
		os.Exit(1)
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Printf("create logger: %v", err)
		os.Exit(1)
	}

	database, err := db.NewDB(
		conf.DbPath,
		"db/migrations",
	)
	if err != nil {
		log.Fatal(err)
	}

	blobStorage, err := storage.NewSqlcRepo(
		database,
		conf.ApiAddr,
		logger,
	)
	if err != nil {
		log.Fatal(err)
	}

	whitelistedPks := make(map[string]struct{})

	api := ginApi.SetupApi(
		blobStorage,
		conf.CdnUrl,
		conf.ApiAddr,
		whitelistedPks,
		logger,
	)
	api.Run()
}
