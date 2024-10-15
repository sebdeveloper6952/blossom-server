package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetSettings(
	ctx context.Context,
	services core.Services,
) ([]*core.Setting, error) {
	return services.Settings().GetAll(ctx)
}
