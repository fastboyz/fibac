package config

import (
	plaidConfig "fibac/config/plaid"
	"fibac/log"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/plaid/plaid-go/v10/plaid"
	"go.uber.org/zap"
)

type IFace interface {
	Log() *zap.Logger
}

type AppConfig struct {
	PlaidApiClient *plaid.APIClient
}

func New() (*AppConfig, error) {
	cfg := &AppConfig{}

	cfg.loadEnv()

	plaidCfg := plaidConfig.Must(plaidConfig.New(cfg))
	cfg.PlaidApiClient = plaidCfg.GetClient()
	return cfg, nil
}

func (c *AppConfig) loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the .env file")
	}
}

func Must(cfg *AppConfig, err error) *AppConfig {
	if err == nil {
		return cfg
	}
	panic(err)
}

func (c *AppConfig) Log() *zap.Logger {
	l := log.Log{}
	return l.Get()
}
