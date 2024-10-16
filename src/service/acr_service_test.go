package service

import (
	"context"
	"testing"

	"github.com/nbd-wtf/go-nostr"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/stretchr/testify/assert"
)

func TestAllowAll(t *testing.T) {
	acr := &acrService{
		rules: map[string][]core.ACR{
			string(core.ResourceUpload): {
				core.ACR{
					Action:   core.ACRActionAllow,
					Pubkey:   "ALL",
					Resource: core.ResourceUpload,
				},
			},
		},
	}

	assert.NoError(
		t,
		acr.Validate(
			context.TODO(),
			"pubkey",
			core.ResourceUpload,
		),
		"ALLOW ALL should return no error",
	)
}

func TestDenyAll(t *testing.T) {
	acr := &acrService{
		rules: map[string][]core.ACR{
			string(core.ResourceUpload): {
				core.ACR{
					Action:   core.ACRActionDeny,
					Pubkey:   "ALL",
					Resource: core.ResourceUpload,
				},
			},
		},
	}

	assert.Error(
		t,
		acr.Validate(
			context.TODO(),
			"pubkey",
			core.ResourceUpload,
		),
		"DENY ALL should return error",
	)
}

func TestAllowAllDenyPubkey(t *testing.T) {
	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())
	acr := &acrService{
		rules: map[string][]core.ACR{
			string(core.ResourceUpload): {
				core.ACR{
					Action:   core.ACRActionAllow,
					Pubkey:   "ALL",
					Resource: core.ResourceUpload,
				},
				core.ACR{
					Action:   core.ACRActionDeny,
					Pubkey:   pk,
					Resource: core.ResourceUpload,
				},
			},
		},
	}

	assert.Error(
		t,
		acr.Validate(
			context.TODO(),
			pk,
			core.ResourceUpload,
		),
		"ALLOW ALL & DENY PK should return error",
	)
}

func TestDenyAllAllowPubkey(t *testing.T) {
	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())
	acr := &acrService{
		rules: map[string][]core.ACR{
			string(core.ResourceUpload): {
				core.ACR{
					Action:   core.ACRActionDeny,
					Pubkey:   "ALL",
					Resource: core.ResourceUpload,
				},
				core.ACR{
					Action:   core.ACRActionAllow,
					Pubkey:   pk,
					Resource: core.ResourceUpload,
				},
			},
		},
	}

	assert.NoError(
		t,
		acr.Validate(
			context.TODO(),
			pk,
			core.ResourceUpload,
		),
		"DENY ALL & ALLOW PK should return nil",
	)
}
