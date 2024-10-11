package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	ginApi "github.com/sebdeveloper6952/blossom-server/api/gin"
	"github.com/sebdeveloper6952/blossom-server/db"
	accesscontrol "github.com/sebdeveloper6952/blossom-server/src/access-control"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/logging"
	"github.com/sebdeveloper6952/blossom-server/src/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conf, err := config.NewConfig("config.yml")
	if err != nil {
		log.Fatalf("new config: %v", err)
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}

	database, err := db.NewDB(
		conf.DbPath,
		"db/migrations",
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("[main][db] %s", err))
	}
	queries := db.New(database)

	blobService, err := service.NewBlobService(
		database,
		queries,
		conf.ApiAddr,
		logger,
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("[main][blob-storage] %s", err))
	}

	acrService, err := service.NewACRService(
		database,
		queries,
		logger,
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("[main][acr-storage] %s", err))
	}

	settingsService, err := service.NewSettingService(
		database,
		queries,
		logger,
	)

	if err := accesscontrol.EnsureAdminHasAccess(
		ctx,
		acrService,
		conf.AdminPubkey,
	); err != nil {
		// TODO: handle error properly
		logger.Error(fmt.Sprintf("[main][ensure-admin-access] %s", err))
	}

	api := ginApi.SetupApi(
		blobService,
		acrService,
		settingsService,
		conf.CdnUrl,
		conf.ApiAddr,
		conf.AdminPubkey,
		conf.UIEnabled,
		logger,
	)
	api.Run()
}
