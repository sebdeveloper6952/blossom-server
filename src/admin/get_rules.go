package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetRules(
	ctx context.Context,
	services core.Services,
) ([]*core.ACR, error) {
	return services.ACR().GetAll(ctx)
}
