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

func CartView(c *gin.Context) {
	if UserID == "" {
		return
	}
	cart := store.ViewCart(UserID)
	c.JSON(200, cart)
}

func Deletecart(c *gin.Context) {
	if UserID == "" {
		return
	}
	var cart []store.CartAdd
	c.ShouldBindJSON(&cart)
	res := store.DeleteProduct(cart, UserID)
	c.JSON(200, res)
}
