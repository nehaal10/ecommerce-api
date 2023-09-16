package server

import (
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func Start(cfg conf.Config) {
	store.SetupDBConeection(cfg)
}
