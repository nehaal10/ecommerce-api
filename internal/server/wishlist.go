package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func WishlistAdd(c *gin.Context) {
	if UserID == "" {
		return
	}
	var payload []store.Wishlist
	c.ShouldBindJSON(&payload)
	result := store.Wishlist_add(payload, UserID)
	c.JSON(200, result)
}

func WishlistView(c *gin.Context) {
	if UserID == "" {
		return
	}

	c.JSON(200, store.ViewWishlist(UserID))
}

func DeleteWishlist(c *gin.Context) {
	if UserID == "" {
		return
	}
	var value []store.Wishlist
	c.ShouldBindJSON(&value)
	response := store.DeleteWishlist(value, UserID)
	c.JSON(200, response)
}
