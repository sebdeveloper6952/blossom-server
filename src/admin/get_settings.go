package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetSettings(
	ctx context.Context,
	settingService core.SettingService,
) ([]*core.Setting, error) {
	return settingService.GetAll(ctx)
}
