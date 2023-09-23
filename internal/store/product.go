package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(prod []Product, id string) {

	var prodInterface []interface{}
	for num, val := range prod {
		val.VendorID = id
		prod[num].VendorID = id
		prodInterface = append(prodInterface, val)
	}
	db.Product.InsertMany(context.TODO(), prodInterface)
	updateVendor(prod, id)
}

func updateVendor(prod []Product, id string) {
	var result VendorDashboard
	filter := bson.M{"admin_id": id}
	res := db.Vendor.FindOne(context.TODO(), filter)
	res.Decode(&result)
	var update primitive.M
	if result.Product == nil {
		update = bson.M{
			"$set": bson.M{
				"product": prod,
			},
		}
	} else {
		update = bson.M{"$push": bson.M{
			"product": bson.M{
				"$each": prod,
			},
		},
		}
	}
	_, err := db.Vendor.UpdateOne(context.TODO(), filter, update)
	utils.Checkerr(err)

}
