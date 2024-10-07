package accesscontrol

import (
	"context"
	"errors"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/src/core"
	"go.uber.org/zap"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrMissingRule  = errors.New("internal server error: missing rule")
)

func Validate(
	ctx context.Context,
	ac core.ACRStorage,
	pubkey string,
	resource core.ACRResource,
	log *zap.Logger,
) error {
	allAcr, err := ac.GetFromPubkeyResource(ctx, "ALL", resource)
	if err != nil {
		// critical error: by core logic, every resource needs to have
		// an "ALL" rule
		log.Error(fmt.Sprintf("[validate] %s", err))
		return ErrMissingRule
	}

	pubkeyAcr, _ := ac.GetFromPubkeyResource(
		ctx,
		pubkey,
		resource,
	)

	return validate(allAcr, pubkeyAcr)
}

func validate(
	allAcr *core.ACR,
	pubkeyAcr *core.ACR,
) error {
	allow := false

	if allAcr != nil {
		allow = allAcr.Action == core.ACRActionAllow
	}

	if pubkeyAcr != nil {
		allow = pubkeyAcr.Action == core.ACRActionAllow
	}

	if allow {
		return nil
	}

	return ErrUnauthorized
}
