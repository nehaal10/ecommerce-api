package store

import (
	"context"
	"net/mail"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type UserLogin struct {
	UserID   string    `json:"user_id,omitempty"`
	EmailID  string    `json:"email_id" validate:"required"`
	Password string    `json:"password" validate:"required,min=4"`
	Add      []Address `json:"address,omitempty"`
}

func (l *UserLogin) EmailValid() bool {
	if _, err := mail.ParseAddress(l.EmailID); err != nil {
		return false
	}
	return true
}

func Login(l UserLogin) (int, string, []Address) {
	if !l.EmailValid() {
		return 400, "CANNOT", nil
	}
	isVal, id, add := getOneUser(l)
	if isVal {
		return 200, id, add
	}
	return 400, "WRONG USER DOESNT EXIST", nil
}

func getOneUser(u UserLogin) (bool, string, []Address) {
	var res User
	filter := bson.M{"email_id": u.EmailID}
	result := db.NewUser.FindOne(context.TODO(), filter)
	result.Decode(&res)
	str := res.UserID
	add := res.Address
	isValid := utils.ComaparePass(res.Password, u.Password)
	return isValid, str, add
}
