package service

import (
	"testing"

	"github.com/nbd-wtf/go-nostr"
	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/stretchr/testify/assert"
)

func TestAllowAll(t *testing.T) {
	allAcr := &core.ACR{
		Action:   core.ACRActionAllow,
		Pubkey:   "ALL",
		Resource: "/resource",
	}

	assert.NoError(
		t,
		validate(
			allAcr,
			nil,
		),
		"ALLOW ALL should return no error",
	)
}

func TestDenyAll(t *testing.T) {
	allAcr := &core.ACR{
		Action:   core.ACRActionDeny,
		Pubkey:   "ALL",
		Resource: "/resource",
	}

	assert.Error(
		t,
		validate(
			allAcr,
			nil,
		),
		"DENY ALL should return error",
	)
}

func TestAllowAllDenyPubkey(t *testing.T) {
	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())
	allAcr := &core.ACR{
		Action:   core.ACRActionAllow,
		Pubkey:   "ALL",
		Resource: "/resource",
	}
	pubkeyAcr := &core.ACR{
		Action:   core.ACRActionDeny,
		Pubkey:   pk,
		Resource: "/resource",
	}

	assert.Error(
		t,
		validate(
			allAcr,
			pubkeyAcr,
		),
		"ALLOW ALL & DENY PK should return error",
	)
}

func TestDenyAllAllowPubkey(t *testing.T) {
	pk, _ := nostr.GetPublicKey(nostr.GeneratePrivateKey())
	allAcr := &core.ACR{
		Action:   core.ACRActionDeny,
		Pubkey:   "ALL",
		Resource: "/resource",
	}
	pubkeyAcr := &core.ACR{
		Action:   core.ACRActionAllow,
		Pubkey:   pk,
		Resource: "/resource",
	}

	assert.NoError(
		t,
		validate(
			allAcr,
			pubkeyAcr,
		),
		"DENY ALL & ALLOW PK should return nil",
	)
}
