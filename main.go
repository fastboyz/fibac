package main

import (
	"fibac/config"
	plaidcfg "fibac/controllers/plaid/config"
	plaidpkg "fibac/controllers/plaid/http"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	cfg := config.Must(config.New())
	logger := cfg.Log()

	router := gin.New()
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginzap.RecoveryWithZap(logger, true))

	router.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	cfgPlaid := plaidcfg.New(cfg)
	ctrlsPlaid := plaidpkg.New(cfgPlaid)
	ctrlsPlaid.Register(router)
	_ = router.Run()
}
