package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/store"
	"github.com/nehaal10/ecommerce-api/internal/utils"
)

type JwtSchema struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func JWT(user store.UserLogin, cfg conf.Config) string {
	expirationTime := time.Now().Add(time.Minute * 2)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtSchema{
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

func GetCoockies(c *gin.Context) (string, error) {
	cfg, err := conf.NewConfig()
	utils.Checkerr(err)
	cook, err := c.Request.Cookie("plswork")
	utils.Checkerr(err)
	token, err := jwt.ParseWithClaims(cook.Value, &JwtSchema{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if err != nil {
		fmt.Println(err)
		utils.Checkerr(err)
	}
	claims, _ := token.Claims.(*JwtSchema)
	UserId := claims.UserID
	return UserId, nil
}
