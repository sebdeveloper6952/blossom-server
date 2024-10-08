package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func DeleteRule(
	ctx context.Context,
	ac core.ACRStorage,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) error {
	return ac.Delete(
		ctx,
		action,
		pubkey,
		resource,
	)
}
