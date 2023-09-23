package server

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

func Start(cfg conf.Config) {
	store.SetupDBConeection(cfg)

	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.POST("/api/auth/register", Register)
	r.POST("/api/auth/login", Login)
	r.POST("/api/admin/register", VendorRegister)
	r.POST("/api/admin/login", Vlogin)
	base_url := "/api/admin/addproduct/"
	apiurl, err := url.Parse(base_url)
	utils.Checkerr(err)
	apiurl.Path = fmt.Sprintf("/vendor/%s", ID)
	api := apiurl.String()
	fmt.Println(api)
	r.POST("/api/admin/addproduct", Product)
	r.Run()
}
