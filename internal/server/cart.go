package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func AddtoCart(c *gin.Context) {
	if UserID == "" {
		return
	}
	var cart []store.CartAdd
	c.ShouldBindJSON(&cart)
	store.UpadateCart(cart, UserID)
}
