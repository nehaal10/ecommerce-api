package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type CustomerDatabase struct {
	UserID          string     `json:"user_id" bson:"user_id"`
	Cart            []CartAdd  `bson:"cart"`
	EmailID         string     `bson:"email_id"`
	Address         []Address  `bson:"Adress"`
	ProductWishList []WishList `bson:"wishlist"`
	Password        string     `bson:"password"`
}

type WishList struct {
	UserId   string    `bson:"user_id"`
	Products []Product `bson:"product"`
}

type Product struct {
	VendorID    string `json:"vendor_id,omitempty" bson:"vendor_id"`
	Category    string `jsno:"category" bson:"category"`
	ProductID   string `jsno:"product_id,omitempty" bson:"product_id"`
	ProductName string `json:"product_name" bson:"product_name"`
	Price       int    `json:"price" bson:"price"`
	Quantity    int    `json:"quantity" bson:"quantity"`
}

func Customer(id string) {
	userid, err := utils.GetCoockies(id)
	utils.Checkerr(err)
	res := getDetailsUser(userid)
	customer := CustomerDatabase{}
	customer.UserID = userid
	customer.EmailID = res.EmailID
	customer.Address = res.Address
	customer.Password = res.Password
	insertToDatabase(customer)
}

func insertToDatabase(cust CustomerDatabase) {
	//check if that user is there
	var result CustomerDatabase
	filter := bson.M{"user_id": cust.UserID}
	res := db.Cutomer.FindOne(context.TODO(), filter)
	res.Decode(&result)
	if result.UserID == "" {
		db.Cutomer.InsertOne(context.TODO(), cust)
	}
}

func getDetailsUser(id string) User {
	var response User
	filter := bson.M{"user_id": id}
	result := db.NewUser.FindOne(context.TODO(), filter)
	err := result.Decode(&response)
	utils.Checkerr(err)
	return response
}
