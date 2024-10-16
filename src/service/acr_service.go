package service

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/sebdeveloper6952/blossom-server/src/core"
	"github.com/sebdeveloper6952/blossom-server/src/pkg/config"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrMissingRule  = errors.New("internal server error: missing rule")
)

type acrService struct {
	rules map[string][]core.ACR
	log   *zap.Logger
}

func NewACRService(
	conf *config.Config,
	log *zap.Logger,
) (core.ACRStorage, error) {
	rules := make(map[string][]core.ACR)
	for _, rule := range conf.AccessControlRules {
		if _, ok := rules[rule.Resource]; !ok {
			rules[rule.Resource] = make([]core.ACR, 0, 2)
		}
		rules[rule.Resource] = append(
			rules[rule.Resource],
			core.ACR{
				Action:   core.ACRAction(rule.Action),
				Pubkey:   rule.Pubkey,
				Resource: core.ACRResource(rule.Resource),
			},
		)
	}

	return &acrService{
		rules: rules,
		log:   log,
	}, nil
}

func (r *acrService) Validate(
	ctx context.Context,
	pubkey string,
	resource core.ACRResource,
) error {
	rules, ok := r.rules[string(resource)]
	if !ok {
		return errors.New("invalid state: there must be at least one rule for the resource")
	}

	allowed := false
	for _, rule := range rules {
		if rule.Pubkey == "ALL" {
			if rule.Action == core.ACRActionAllow {
				allowed = true
			} else {
				allowed = false
			}
		}

		if rule.Pubkey == pubkey {
			if rule.Action == core.ACRActionAllow {
				allowed = true
			} else {
				allowed = false
			}
			break
		}
	}

	if !allowed {
		return ErrUnauthorized
	}

	return nil
}
