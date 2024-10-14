package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func UpdateSetting(
	ctx context.Context,
	settingService core.SettingService,
	key string,
	value string,
) (*core.Setting, error) {
	return settingService.Update(ctx, key, value)
}
