package store

import (
	"context"
	"os/exec"
	"strings"

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

type Vendor struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	First_Name string             `json:"first_name" bson:"first_name"`
	Last_Name  string             `json:"last_name" bson:"last_name"`
	EmailID    string             `json:"email_id" bson:"email_id"`
	Company    string             `json:"company" bson:"company"`
	Password   string             `json:"password" bson:"password" validate:"required,min=4"`
	Phone_No   string             `json:"phone_no" bson:"phone_no"`
	Key        string             `json:"key" bson:"key" validate:"required"`
	Admin_id   string             `json:"admin_id,omitempty" bson:"admin_id"`
}

type Address struct {
	House   string `json:"house_name" bson:"house_name"`
	Street  string `json:"street_name" bson:"street_name"`
	City    string `json:"city_name" bson:"city_name"`
	Pincode string `json:"pin_code" bson:"pin_code"`
}

func AddUser(user User) string {
	user.UserID = GenerateUniqueID()
	user.Password = utils.PasswordHash(user.Password)
	if checkEmail(user.EmailID, user.Phone_No) {
		db.NewUser.InsertOne(context.TODO(), user)
		return "USER ADDED"
	}
	return "DUPLICATE DETECTED"
}

// checks if that mail or phone number exits if it does then the data wont be added to the database
func checkEmail(email string, num string) bool {
	var returned []User
	filter := bson.M{
		"$or": []bson.M{
			{"email_id": email},
			{"phone_no": num},
		},
	}
	result, _ := db.NewUser.Find(context.TODO(), filter)
	err := result.All(context.TODO(), &returned)
	utils.Checkerr(err)
	return len(returned) == 0
}

func checkVendor(email string, num string) bool {
	var returned []Vendor
	filter := bson.M{
		"$or": []bson.M{
			{"email_id": email},
			{"phone_no": num},
		},
	}
	result, _ := db.VendorRegister.Find(context.TODO(), filter)
	err := result.All(context.TODO(), &returned)
	utils.Checkerr(err)
	return len(returned) == 0
}

func GenerateUniqueID() string {
	id := ksuid.New()
	return id.String()
}

func uniqueKey() string {
	newUUID, err := exec.Command("uuidgen").Output()
	utils.Checkerr(err)
	str := string(newUUID)
	str = strings.TrimSpace(str)
	return str
}

func RegisterAdmin(vendor Vendor) (int, string) {
	vendor.Admin_id = GenerateUniqueID()
	vendor.Password = utils.PasswordHash(vendor.Password)
	//assigning unique key for the vendor
	vendor.Key = uniqueKey()
	if checkVendor(vendor.EmailID, vendor.Phone_No) {
		db.VendorRegister.InsertOne(context.TODO(), vendor)
		return 200, "VENDOR ADDED"
	}
	//db.Vendor.InsertOne()
	return 400, "DUPLICATE"
}
