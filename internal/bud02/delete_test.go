package bud02

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nbd-wtf/go-nostr"
	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/db"
	"github.com/sebdeveloper6952/blossom-server/internal/core"
	"github.com/sebdeveloper6952/blossom-server/internal/pkg/config"
	"github.com/sebdeveloper6952/blossom-server/internal/pkg/hashing"
	"github.com/sebdeveloper6952/blossom-server/internal/pkg/logging"
	"github.com/sebdeveloper6952/blossom-server/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBlob(t *testing.T) {
	dbFile := "./db-TestDeleteBlob.sqlite3"
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
		AccessControlRules: []config.AccessControlRule{
			{Action: string(core.ACRActionAllow), Pubkey: "ALL", Resource: string(core.ResourceUpload)},
		},
		AllowedMimeTypes: []string{"*"},
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}

	database, err := db.NewDB(dbFile, "../../db/migrations")
	if err != nil {
		t.Fatal(err)
	}
	queries := db.New(database)

	services := service.New(context.TODO(), database, queries, conf, logger)
	services.Init(context.TODO())

	blobBytes := []byte{}
	authHash, _ := hashing.Hash(blobBytes)

	// upload blob first
	_, err = UploadBlob(
		context.TODO(),
		services,
		conf.CdnUrl,
		authHash,
		pk,
		blobBytes,
	)
	assert.NoError(t, err, "upload should succeed")

	// delete the blob
	err = DeleteBlob(
		context.TODO(),
		services,
		pk,
		authHash,
		authHash,
		zap.NewNop(),
	)
	assert.NoError(t, err, "delete should succeed")

	// verify blob no longer exists
	_, err = services.Blob().GetFromHash(context.TODO(), authHash)
	assert.Error(t, err, "blob should not exist after deletion")
}

func TestDeleteBlobNotOwner(t *testing.T) {
	dbFile := "./db-TestDeleteBlobNotOwner.sqlite3"
	defer func() {
		if err := os.Remove(dbFile); err != nil {
			t.Log(err)
		}
	}()

	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())
	otherPk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())

	conf := &config.Config{
		DbPath:      dbFile,
		LogLevel:    "DEBUG",
		CdnUrl:      "http://localhost:8000",
		AdminPubkey: pk,
		AccessControlRules: []config.AccessControlRule{
			{Action: string(core.ACRActionAllow), Pubkey: "ALL", Resource: string(core.ResourceUpload)},
		},
		AllowedMimeTypes: []string{"*"},
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}

	database, err := db.NewDB(dbFile, "../../db/migrations")
	if err != nil {
		t.Fatal(err)
	}
	queries := db.New(database)

	services := service.New(context.TODO(), database, queries, conf, logger)
	services.Init(context.TODO())

	blobBytes := []byte{}
	authHash, _ := hashing.Hash(blobBytes)

	// upload blob as pk
	_, err = UploadBlob(
		context.TODO(),
		services,
		conf.CdnUrl,
		authHash,
		pk,
		blobBytes,
	)
	assert.NoError(t, err, "upload should succeed")

	// try to delete as a different pubkey
	err = DeleteBlob(
		context.TODO(),
		services,
		otherPk,
		authHash,
		authHash,
		zap.NewNop(),
	)
	assert.Error(t, err, "delete by non-owner should fail")
}

func TestDeleteBlobNotFound(t *testing.T) {
	dbFile := "./db-TestDeleteBlobNotFound.sqlite3"
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
		AccessControlRules: []config.AccessControlRule{
			{Action: string(core.ACRActionAllow), Pubkey: "ALL", Resource: string(core.ResourceUpload)},
		},
		AllowedMimeTypes: []string{"*"},
	}

	logger, err := logging.NewLog(conf.LogLevel)
	if err != nil {
		log.Fatalf("new logger: %v", err)
	}

	database, err := db.NewDB(dbFile, "../../db/migrations")
	if err != nil {
		t.Fatal(err)
	}
	queries := db.New(database)

	services := service.New(context.TODO(), database, queries, conf, logger)
	services.Init(context.TODO())

	// try to delete a blob that was never uploaded
	err = DeleteBlob(
		context.TODO(),
		services,
		pk,
		"nonexistenthash",
		"nonexistenthash",
		zap.NewNop(),
	)
	assert.ErrorIs(t, err, core.ErrBlobNotFound, "should return blob not found error")
}
