package bud04

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nbd-wtf/go-nostr"
	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/hashing"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/logging"
	"github.com/sebdeveloper6952/blossom-server/src/service"
	"github.com/stretchr/testify/assert"
)

func TestMirrorUnauth(t *testing.T) {
	dbFile := "./db-TestMirrorUnauth.sqlite3"
	defer func() {
		if err := os.Remove(dbFile); err != nil {
			t.Log(err)
		}
	}()

	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())

	conf := &config.Config{
		DbPath:      dbFile,
		LogLevel:    "DEBUG",
		CdnUrl:      "http://localhost:8000",
		AdminPubkey: pk,
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}

	database, err := db.NewDB(
		dbFile,
		"../../db/migrations",
	)
	if err != nil {
		t.Fatal(err)
	}
	queries := db.New(database)

	services := service.New(database, queries, conf, logger)

	blobBytes := make([]byte, 32)
	authHash, _ := hashing.Hash(blobBytes)
	blobURL := url.URL{}

	_, err = MirrorBlob(
		context.TODO(),
		services,
		conf.CdnUrl,
		pk,
		authHash,
		blobURL,
	)

	assert.Error(t, err, "expected unauthorized error")
}
