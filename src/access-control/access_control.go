package accesscontrol

import (
	"context"
	"errors"
	"fmt"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func Validate(
	ctx context.Context,
	ac core.ACRStorage,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) error {
	allAcr, err := ac.GetFromPubkeyResource(ctx, "ALL", resource)
	if err != nil {
		// critical error: by core logic, every resource needs to have
		// an "ALL" rule
		fmt.Println(err)
	}

	pubkeyAcr, err := ac.GetFromPubkeyResource(
		ctx,
		pubkey,
		resource,
	)
	if err != nil {
		fmt.Println(err)
	}

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

	return errors.New("unauthorized")
}
