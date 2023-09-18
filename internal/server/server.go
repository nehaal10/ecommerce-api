package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func Start(cfg conf.Config) {
	store.SetupDBConeection(cfg)

	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.POST("/api/auth/register", Register)
	r.POST("/api/auth/login", Login)
	r.Run()
}
