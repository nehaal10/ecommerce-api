package server

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

type JwtSchema struct {
	UserID string
	jwt.RegisteredClaims
}

func JWT(user store.UserLogin, cfg conf.Config) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtSchema{
		UserID: user.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.UserID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
		},
	})
	token, err := claims.SignedString([]byte(cfg.Secret))
	utils.Checkerr(err)
	return token
}
