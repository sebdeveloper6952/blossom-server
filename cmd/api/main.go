package main

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"

	ginApi "github.com/sebdeveloper6952/blossom-server/v2/internal/httpapi"
	"github.com/sebdeveloper6952/blossom-server/v2/db"
	"github.com/sebdeveloper6952/blossom-server/v2/internal/pkg/config"
	"github.com/sebdeveloper6952/blossom-server/v2/internal/pkg/logging"
	"github.com/sebdeveloper6952/blossom-server/v2/internal/service"
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
		logger.Fatal(err.Error())
	}
	queries := db.New(database)

	services := service.New(
		ctx,
		database,
		queries,
		conf,
		logger,
	)
	if err := services.Init(ctx); err != nil {
		logger.Error(err.Error())
	}

	api := ginApi.SetupRoutes(
		services,
		conf.CdnUrl,
		conf.AdminPubkey,
		logger,
	)
	api.Run(conf.ApiAddr)
}
