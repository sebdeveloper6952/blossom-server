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
	services core.Services,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) (*core.ACR, error) {
	if !nostr.IsValidPubkey(pubkey) {
		return nil, ErrInvalidPubkey
	}

	return services.ACR().Save(
		ctx,
		action,
		pubkey,
		resource,
	)
}
