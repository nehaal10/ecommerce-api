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
	r.POST("/api/admin/register", VendorRegister)
	r.POST("/api/admin/login", Vlogin)
	r.POST("/api/admin/addproduct", Product)
	r.GET("/api/view/products", ShowAllProduct)
	r.POST("/api/user/cart/add", AddtoCart)
	r.GET("/api/user/cart/view", CartView)
	r.DELETE("/api/user/cart/remove", Deletecart)
	r.POST("/api/user/wishlist/add", WishlistAdd)
	r.GET("/api/user/wishlist/view", WishlistView)
	r.DELETE("/api/user/wishlist/remove", DeleteWishlist)
	r.Run()
}
