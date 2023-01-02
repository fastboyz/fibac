package plaid

import (
	svcfg "fibac/controllers/plaid/config"
	"fibac/controllers/plaid/http/handlers"
	"github.com/gin-gonic/gin"
)

const (
	getAccessTokenPath = "/api/set_access_token"
)

type Service struct {
	cfg svcfg.IFace
}

func New(cfg svcfg.IFace) (s *Service) {
	return &Service{
		cfg: cfg,
	}
}

func (s *Service) GetAccessToken() gin.HandlerFunc {
	return handlers.GetAccessTokenHandler(s.cfg)
}

func (s *Service) Register(engine *gin.Engine) *gin.Engine {
	engine.POST(getAccessTokenPath, s.GetAccessToken())
	return engine
}
