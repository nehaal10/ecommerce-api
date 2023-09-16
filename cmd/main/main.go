package main

import (
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/server"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

func main() {
	conf, err := conf.NewConfig()
	utils.Checkerr(err)
	server.Start(conf)
}
