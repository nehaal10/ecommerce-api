package store

import (
	"context"
	"net/mail"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserLogin struct {
	UserID   string `json:"user_id,omitempty"`
	EmailID  string `json:"email_id" validate:"required"`
	Password string `json:"password" validate:"required,min=4"`
}

func (l *UserLogin) EmailValid() bool {
	if _, err := mail.ParseAddress(l.EmailID); err != nil {
		return false
	}
	return true
}

func Login(l UserLogin) (int, string) {
	if !l.EmailValid() {
		return 400, "CANNOT"
	}
	isVal, id := getOneUser(l)
	if isVal {
		return 200, id
	}
	return 400, "WRONG USER DOESNT EXIST"
}

func getOneUser(u UserLogin) (bool, string) {
	var res User
	filter := bson.M{"email_id": u.EmailID}
	result := db.NewUser.FindOne(context.TODO(), filter)
	result.Decode(&res)
	str := res.UserID
	isValid := utils.ComaparePass(res.Password, u.Password)
	return isValid, str
}
