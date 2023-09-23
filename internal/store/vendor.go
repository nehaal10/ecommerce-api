package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type VendorDashboard struct {
	Vendor_id string    `bson:"admin_id"`
	EmailID   string    `bson:"email_id"`
	Password  string    `bson:"password"`
	Key       string    `bson:"key"`
	Product   []Product `bson:"product"`
}

func VendorGet(token string) {
	id, err := utils.GetCoockiesVendor(token)
	utils.Checkerr(err)
	v := VendorDashboard{}
	details := getOnevendor(id)
	v.Vendor_id = id
	v.EmailID = details.EmailID
	v.Password = details.Password
	v.Key = details.Key
	insertToVendorDatabase(v)
	//utils.Checkerr(err)
}

func insertToVendorDatabase(v VendorDashboard) {
	var result VendorDashboard
	filter := bson.M{"admin_id": v.Vendor_id}
	res := db.Vendor.FindOne(context.TODO(), filter)
	res.Decode(&result)
	if result.Vendor_id == "" {
		db.Vendor.InsertOne(context.TODO(), v)
		// return nil
	}

	//return errors.New("already in the database")
}

func getOnevendor(id string) Vendor {
	var result Vendor
	filter := bson.M{"admin_id": id}
	res := db.VendorRegister.FindOne(context.TODO(), filter)
	err := res.Decode(&result)
	utils.Checkerr(err)
	return result
}
