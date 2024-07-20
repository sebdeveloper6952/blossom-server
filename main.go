package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"

	ginApi "github.com/sebdeveloper6952/blossom-server/api/gin"
	"github.com/sebdeveloper6952/blossom-server/db"
	blobDescriptorRepos "github.com/sebdeveloper6952/blossom-server/repos/blob_descriptor"
)

var (
	serviceName  = "blossom-server"
	collectorURL = "localhost:4317"
	insecure     = "true"
)

func initTracer() func(context.Context) error {
	var secureOption otlptracegrpc.Option

	if strings.ToLower(insecure) == "false" || insecure == "0" || strings.ToLower(insecure) == "f" {
		secureOption = otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	} else {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Fatalf("Failed to create exporter: %v", err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Fatalf("Could not set resources: %v", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}

func main() {
	traceCleanup := initTracer()
	defer traceCleanup(context.Background())

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
		serviceName,
		blobDescriptorRepo,
		config.CdnUrl,
		config.ApiAddr,
		whitelistedPks,
		logger,
	)
	api.Run()
}
