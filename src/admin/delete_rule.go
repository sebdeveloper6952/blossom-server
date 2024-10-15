package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func DeleteRule(
	ctx context.Context,
	services core.Services,
	action core.ACRAction,
	pubkey string,
	resource core.ACRResource,
) error {
	return services.ACR().Delete(
		ctx,
		action,
		pubkey,
		resource,
	)
}
