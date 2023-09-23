package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func Product(c *gin.Context) {
	if ID == "" {
		return
	}
	var prod []store.Product
	c.ShouldBindJSON(&prod)
	store.AddProduct(prod, ID)
}
