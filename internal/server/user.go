package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

func Register(c *gin.Context) {
	var user store.User
	err := c.ShouldBindJSON(&user)
	utils.Checkerr(err)
	result := store.AddUser(user)
	c.JSON(200, gin.H{
		"message": result,
	})
}

func Login(c *gin.Context) {
	var userlogin store.UserLogin
	err := c.ShouldBindJSON(&userlogin)
	utils.Checkerr(err)
	num, res := store.Login(userlogin)

	if num != 200 {
		c.JSON(num, gin.H{
			"Message": res,
		})
		return
	}
	userlogin.UserID = res
	cfg, err := conf.NewConfig()
	utils.Checkerr(err)
	token := JWT(userlogin, cfg)
	c.JSON(num, gin.H{
		"token": token,
	})
}
