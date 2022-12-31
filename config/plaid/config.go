package plaid

import (
	"errors"
	appcfg "fibac/config"
	"github.com/plaid/plaid-go/v10/plaid"
	"os"
)

var environments = map[string]plaid.Environment{
	"sandbox":     plaid.Sandbox,
	"development": plaid.Development,
	"production":  plaid.Production,
}

type IFace interface {
	GetClient() *plaid.APIClient
}

type Config struct {
	appcfg.IFace
	ClientId     string
	Secret       string
	Env          string
	Products     string
	CountryCodes string
	RedirectUri  string
	client       *plaid.APIClient
}

func New(base appcfg.IFace) (*Config, error) {
	cfg := &Config{
		IFace:        base,
		ClientId:     os.Getenv("PLAID_CLIENT_ID"),
		Secret:       os.Getenv("PLAID_SECRET"),
		Env:          os.Getenv("PLAID_ENV"),
		Products:     os.Getenv("PLAID_PRODUCTS"),
		CountryCodes: os.Getenv("PLAID_COUNTRY_CODES"),
		RedirectUri:  os.Getenv("PLAID_REDIRECT_URI"),
		client:       nil,
	}

	initDefaults(cfg)

	if cfg.ClientId == "" || cfg.Secret == "" {
		return nil, errors.New("error: PLAID_SECRET or PLAID_CLIENT_ID was not set")
	}

	conf := plaid.NewConfiguration()
	conf.AddDefaultHeader("PLAID-CLIENT-ID", cfg.ClientId)
	conf.AddDefaultHeader("PLAID-SECRET", cfg.Secret)
	conf.UseEnvironment(environments[cfg.Env])
	cfg.client = plaid.NewAPIClient(conf)

	return cfg, nil
}

func initDefaults(cfg *Config) {
	if cfg.Products == "" {
		cfg.Products = " transactions"
	}
	if cfg.CountryCodes == "" {
		cfg.CountryCodes = "CA"
	}
	if cfg.Env == "" {
		cfg.Env = "sandbox"
	}
}

func (cfg *Config) GetClient() *plaid.APIClient {
	return cfg.client
}

func Must(cfg *Config, err error) *Config {
	if err == nil {
		return cfg
	}
	panic(err)
}
