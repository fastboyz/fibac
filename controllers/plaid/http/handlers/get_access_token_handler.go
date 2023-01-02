package handlers

import (
	"context"
	"fibac/controllers/plaid/config"
	"github.com/gin-gonic/gin"
	"github.com/plaid/plaid-go/v10/plaid"
	"go.uber.org/zap"
	"net/http"
)

func GetAccessTokenHandler(cfg config.IFace) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		publicToken := ginCtx.PostForm("public_token")
		ctx := context.Background()

		resp, err := cfg.PlaidService().PublicTokenExchange(ctx, publicToken)
		if err != nil {
			cfg.Log().Error("Failed Get Access token", zap.Error(err))
			RenderError(ginCtx, err)
		}
		accessToken := resp.GetAccessToken()
		itemID := resp.GetItemId()

		ginCtx.JSON(http.StatusOK, gin.H{
			"access_token": accessToken,
			"item_id":      itemID,
		})
	}
}

func RenderError(c *gin.Context, originalErr error) {
	if plaidError, err := plaid.ToPlaidError(originalErr); err == nil {
		c.JSON(http.StatusOK, gin.H{"error": plaidError})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": originalErr.Error()})
}
