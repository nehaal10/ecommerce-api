package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"github.com/segmentio/ksuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	First_Name string             `json:"first_name" bson:"first_name" validate:"required"`
	Last_Name  string             `json:"last_name" bson:"last_name" validate:"required"`
	Password   string             `json:"password" bson:"password" validate:"required,min=4"`
	EmailID    string             `json:"email_id" bson:"email_id" validate:"required"`
	Phone_No   string             `json:"phone_no" bson:"phone_no" validate:"required"`
	UserID     string             `json:"user_id,omitempty" bson:"user_id"`
	Address    []Address          `json:"address,omitempty" bson:"address" validate:"required"`
}

type Address struct {
	House   string `json:"house_name" bson:"house_name"`
	Street  string `json:"street_name" bson:"street_name"`
	City    string `json:"city_name" bson:"city_name"`
	Pincode string `json:"pin_code" bson:"pin_code"`
}

func AddUser(user User) string {
	user.UserID = generateUniqueID()
	user.Password = utils.PasswordHash(user.Password)
	if checkEmail(user.EmailID, user.Phone_No) {
		db.NewUser.InsertOne(context.TODO(), user)
		return "USER ADDED"
	}
	return "DUPLICATE DETECTED"
}

func checkEmail(email string, num string) bool {
	var returned []User
	filter := bson.M{
		"$and": []bson.M{
			{"email_id": email},
			{"phone_no": num},
		},
	}
	result, _ := db.NewUser.Find(context.TODO(), filter)
	err := result.All(context.TODO(), &returned)
	utils.Checkerr(err)
	return len(returned) == 0
}

func generateUniqueID() string {
	id := ksuid.New()
	return id.String()
}
