package plaid

import (
	"context"
	"fibac/config/settings"
	"github.com/plaid/plaid-go/v10/plaid"
	"go.uber.org/zap"
	"net/http"
)

type plClient struct {
	log      *zap.Logger
	settings *settings.Settings
	plaid    *plaid.APIClient
}

//go:generate mockery --name PlaidClient --inpackage --case snake
type PlaidClient interface {
	ItemPublicTokenExchangeRequest(ctx context.Context, publicToken string) (plaid.ItemPublicTokenExchangeResponse, *http.Response, error)
}

func NewClient(log *zap.Logger, settings *settings.Settings, plaidClient *plaid.APIClient) PlaidClient {
	return &plClient{
		log:      log,
		settings: settings,
		plaid:    plaidClient,
	}
}

func (pc *plClient) ItemPublicTokenExchangeRequest(ctx context.Context, publicToken string) (plaid.ItemPublicTokenExchangeResponse, *http.Response, error) {
	return pc.plaid.PlaidApi.ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(
		*plaid.NewItemPublicTokenExchangeRequest(publicToken)).Execute()
}
