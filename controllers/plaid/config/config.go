package config

import (
	appcfg "fibac/config"
	services "fibac/services/plaid"
)

type IFace interface {
	appcfg.IFace
	PlaidService() services.PlaidService
}

type PlaidConfig struct {
	appcfg.IFace
	plaidService services.PlaidService
}

func (cfg *PlaidConfig) PlaidService() services.PlaidService {
	return cfg.plaidService
}

func New(base appcfg.IFace) *PlaidConfig {
	cfg := &PlaidConfig{
		IFace:        base,
		plaidService: services.NewPlaidService(base, base.Log()),
	}
	return cfg
}

var _ IFace = &PlaidConfig{}
