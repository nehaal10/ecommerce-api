package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/store"
)

func Product(c *gin.Context) {
	if Vendor_id == "" {
		return
	}
	var prod []store.Product
	c.ShouldBindJSON(&prod)
	store.AddProduct(prod, Vendor_id)
}

func ShowAllProduct(c *gin.Context) {
	queryParam := c.Query("name")
	queryCategory := c.Query("category")
	if queryParam == "" && queryCategory == "" {
		products := store.AllProducts()
		c.JSON(200, products)
		return
	} else if queryCategory == "" {
		product := searchProduct(queryParam)
		if product == nil {
			c.JSON(200, gin.H{
				"message": "No Products",
			})
			return
		}
		c.JSON(200, product)
	} else {
		product := searchCategory(queryCategory)
		if product == nil {
			c.JSON(200, gin.H{
				"message": "No Products",
			})
			return
		}
		c.JSON(200, product)
	}
}

func searchProduct(query string) []store.ViewProduct {
	products := store.SearchSpecificProduct(query)
	return products
}
func searchCategory(query string) []store.ViewProduct {
	prod := store.SearchCategory(query)
	return prod
}
