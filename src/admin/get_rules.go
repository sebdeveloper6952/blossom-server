package admin

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func GetRules(
	ctx context.Context,
	ac core.ACRStorage,
) ([]*core.ACR, error) {
	return ac.GetAll(ctx)
}
