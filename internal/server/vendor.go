package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

var Vendor_id string

func VendorRegister(c *gin.Context) {
	var v store.Vendor
	err := c.ShouldBindJSON(&v)
	utils.Checkerr(err)
	num, str := store.RegisterAdmin(v)
	c.JSON(num, str)

}

func Vlogin(c *gin.Context) {
	cnf, err := conf.NewConfig()
	utils.Checkerr(err)
	var vl store.VendorLogin
	err = c.ShouldBindJSON(&vl)
	utils.Checkerr(err)
	num, str := store.LoginVendor(vl)
	Vendor_id = str
	if num != 200 {
		c.JSON(num, gin.H{
			"Message": str,
		})
		return
	}
	vl.Admin_id = str
	Token := JWTVendor(vl, cnf)
	c.SetCookie("vendor", Token, int(time.Now().Add(time.Minute*2).Unix()), "/", "localhost", true, true)
	store.VendorGet(Token)
	c.JSON(num, gin.H{
		"Message": "Corrrect Login",
	})
}
