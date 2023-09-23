package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nehaal10/ecommerce-api/internal/conf"
	"golang.org/x/crypto/bcrypt"
)

type jwtSchema struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type vendorJWTschema struct {
	Admin_id string `json:"admin_id"`
	jwt.RegisteredClaims
}

func Checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PasswordHash(password string) string {
	datapass, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	Checkerr(err)

	return string(datapass)
}

func ComaparePass(hashed string, pass string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass)); err != nil {
		return false
	}

	return true

}

func GetCoockies(val string) (string, error) {
	cfg, err := conf.NewConfig()
	Checkerr(err)
	token, err := jwt.ParseWithClaims(val, &jwtSchema{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		fmt.Println(err)
		Checkerr(err)
	}
	claims, _ := token.Claims.(*jwtSchema)
	UserId := claims.UserID
	return UserId, nil
}

func GetCoockiesVendor(val string) (string, error) {
	cfg, err := conf.NewConfig()
	Checkerr(err)
	token, err := jwt.ParseWithClaims(val, &vendorJWTschema{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})

	if err != nil {
		fmt.Println(err)
		Checkerr(err)
	}
	claims, _ := token.Claims.(*vendorJWTschema)
	UserId := claims.Admin_id
	return UserId, nil
}
