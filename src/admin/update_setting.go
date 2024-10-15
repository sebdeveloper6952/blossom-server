package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func UpdateSetting(
	ctx context.Context,
	services core.Services,
	key string,
	value string,
) (*core.Setting, error) {
	return services.Settings().Update(ctx, key, value)
}
