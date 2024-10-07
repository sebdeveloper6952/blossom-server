package accesscontrol

import (
	"context"

	"github.com/sebdeveloper6952/blossom-server/src/core"
)

func EnsureAdminHasAccess(
	ctx context.Context,
	ac core.ACRStorage,
	pubkey string,
) error {
	rules := []*core.ACR{
		{
			Action:   core.ACRActionAllow,
			Pubkey:   pubkey,
			Resource: core.ResourceUpload,
		},
		{
			Action:   core.ACRActionAllow,
			Pubkey:   pubkey,
			Resource: core.ResourceGet,
		},
		{
			Action:   core.ACRActionAllow,
			Pubkey:   pubkey,
			Resource: core.ResourceList,
		},
		{
			Action:   core.ACRActionAllow,
			Pubkey:   pubkey,
			Resource: core.ResourceDelete,
		},
		{
			Action:   core.ACRActionAllow,
			Pubkey:   pubkey,
			Resource: core.ResourceMirror,
		},
	}

	return ac.SaveMany(ctx, rules)
}
