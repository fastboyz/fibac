package config

import (
	plaidClient "fibac/clients/plaid"
	"fibac/config/settings"
	"fibac/log"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/plaid/plaid-go/v10/plaid"
	"go.uber.org/zap"
)

const namespace = "fibac"

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

type IFace interface {
	Log() *zap.Logger
	Settings() *settings.Settings
	PlaidClient() plaidClient.PlaidClient
}

type AppConfig struct {
	settings    settings.Settings
	plaidClient plaidClient.PlaidClient
	plClient    *plaid.APIClient
}

func New() (*AppConfig, error) {
	cfg := &AppConfig{}

	cfg.loadEnv()
	cfg.initPlaidClient()
	return cfg, nil
}

func (c *AppConfig) loadEnv() {
	// Loading the .env if one does exist
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading the .env file")
	}

	if err := envconfig.Process(namespace, &c.settings); err != nil {
		c.Log().Fatal(err.Error())
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

func (c *AppConfig) initPlaidClient() {
	if c.settings.ClientId == "" || c.settings.Secret == "" {
		c.Log().Fatal("Error: PLAID_SECRET or PLAID_CLIENT_ID is not set.")
	}
	plaidCfg := plaid.NewConfiguration()
	plaidCfg.AddDefaultHeader("PLAID-CLIENT-ID", c.settings.ClientId)
	plaidCfg.AddDefaultHeader("PLAID-SECRET", c.settings.Secret)
	plaidCfg.UseEnvironment(environments[c.settings.Env])
	client := plaid.NewAPIClient(plaidCfg)
	c.plClient = client

	c.plaidClient = plaidClient.NewClient(c.Log(), &c.settings, c.plClient)
}

func (c *AppConfig) Settings() *settings.Settings {
	return &c.settings
}

func (c *AppConfig) PlaidClient() plaidClient.PlaidClient {
	return c.plaidClient
}
