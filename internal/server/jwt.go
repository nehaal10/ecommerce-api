package server

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

type jwtSchema struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type vendorJWTschema struct {
	Admin_id string `json:"admin_id"`
	jwt.RegisteredClaims
}

func JWT(user store.UserLogin, cfg conf.Config) string {
	expirationTime := time.Now().Add(time.Minute * 2)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtSchema{
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.UserID,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})
	token, err := claims.SignedString([]byte(cfg.Secret))
	utils.Checkerr(err)
	return token
}

func JWTVendor(ven store.VendorLogin, cfg conf.Config) string {
	expirationTime := time.Now().Add(time.Minute * 2)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, vendorJWTschema{
		Admin_id: ven.Admin_id,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    ven.Admin_id,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	})
	token, err := claims.SignedString([]byte(cfg.Secret))
	utils.Checkerr(err)
	return token
}
