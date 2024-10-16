package core

import (
	"context"
)

type ACRAction string

const (
	ACRActionAllow ACRAction = "ALLOW"
	ACRActionDeny  ACRAction = "DENY"
)

type ACRResource string

const (
	ResourceUpload ACRResource = "UPLOAD"
	ResourceGet    ACRResource = "GET"
	ResourceDelete ACRResource = "DELETE"
	ResourceList   ACRResource = "LIST"
	ResourceMirror ACRResource = "MIRROR"
)

type ACR struct {
	Action   ACRAction
	Pubkey   string
	Resource ACRResource
}

func NewACR(action ACRAction, pubkey string, resource ACRResource) *ACR {
	return &ACR{
		Action:   action,
		Pubkey:   pubkey,
		Resource: resource,
	}
}

type ACRStorage interface {
	Validate(
		ctx context.Context,
		pubkey string,
		resource ACRResource,
	) error
}
