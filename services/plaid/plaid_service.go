package plaid

import (
	"context"
	"fibac/config"
	"github.com/plaid/plaid-go/v10/plaid"
	"go.uber.org/zap"
)

//go:generate mockery --name PlaidService --inpackage --case snake
type PlaidService interface {
	PublicTokenExchange(ctx context.Context, publicToken string) (*plaid.ItemPublicTokenExchangeResponse, error)
}

type plaidService struct {
	cfg config.IFace
	log *zap.Logger
}

func NewPlaidService(cfg config.IFace, log *zap.Logger) PlaidService {
	return &plaidService{
		cfg: cfg,
		log: log,
	}
}

func (p *plaidService) PublicTokenExchange(ctx context.Context, publicToken string) (*plaid.ItemPublicTokenExchangeResponse, error) {
	publicTokenExchangeResponse, _, err := p.cfg.PlaidClient().ItemPublicTokenExchangeRequest(ctx, publicToken)
	if err != nil {
		return nil, err
	}
	return &publicTokenExchangeResponse, nil
}
