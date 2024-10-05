package core

import "context"

type ACRAction string

const (
	ACRActionAllow ACRAction = "ALLOW"
	ACRActionDeny  ACRAction = "DENY"
)

type ACRResource string

const (
	ResourceUpload ACRResource = "UPLOAD"
)

type ACR struct {
	Action   ACRAction
	Pubkey   string
	Resource ACRResource
	Priority int
}

func NewACR(action ACRAction, pubkey string, resource ACRResource) *ACR {
	return &ACR{
		Action:   action,
		Pubkey:   pubkey,
		Resource: resource,
	}
}

type ACRStorage interface {
	Save(
		ctx context.Context,
		action ACRAction,
		pubkey string,
		resource ACRResource,
		priority int,
	) (*ACR, error)
	Get(
		ctx context.Context,
		action ACRAction,
		pubkey string,
		resource ACRResource,
	) (*ACR, error)
	GetFromPubkey(
		ctx context.Context,
		pubkey string,
	) ([]*ACR, error)
	GetFromPubkeyResource(
		ctx context.Context,
		pubkey string,
		resource ACRResource,
	) (*ACR, error)
	Delete(
		ctx context.Context,
		action ACRAction,
		pubkey string,
		resource ACRResource,
	) error
}
