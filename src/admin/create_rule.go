package admin

import (
	"context"
	"errors"

	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/nostr"
)

var (
	ErrInvalidPubkey = errors.New("invalid pubkey")
)

func CreateRule(
	ctx context.Context,
	ac core.ACRStorage,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) (*core.ACR, error) {
	if !nostr.IsValidPubkey(pubkey) {
		return nil, ErrInvalidPubkey
	}

	return ac.Save(
		ctx,
		action,
		pubkey,
		resource,
	)
}
