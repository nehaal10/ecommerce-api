package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ViewProduct struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
	Quantity    int    `json:"quantity"`
}

func AddProduct(prod []Product, id string) {

	var prodInterface []interface{}
	for num, val := range prod {
		prod[num].VendorID = id
		prod[num].ProductID = GenerateUniqueID()
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

func AllProducts() []ViewProduct {
	var products []ViewProduct
	filter := bson.M{}
	curr, err := db.Product.Find(context.TODO(), filter)
	utils.Checkerr(err)
	for curr.Next(context.TODO()) {
		var product Product
		var prod ViewProduct
		curr.Decode(&product)
		prod.ProductName = product.ProductName
		prod.Price = product.Price
		prod.Quantity = product.Quantity
		products = append(products, prod)
	}

	return products
}

func SearchSpecificProduct(productname string) []ViewProduct {
	var products []ViewProduct
	filter := bson.M{"product_name": productname}
	curr, err := db.Product.Find(context.TODO(), filter)
	utils.Checkerr(err)

	for curr.Next(context.TODO()) {
		var product Product
		var prod ViewProduct
		curr.Decode(&product)
		prod.ProductName = product.ProductName
		prod.Price = product.Price
		prod.Quantity = product.Quantity
		products = append(products, prod)
	}
	return products
}

func SearchCategory(catName string) []ViewProduct {
	var products []ViewProduct
	filter := bson.M{"category": catName}
	curr, err := db.Product.Find(context.TODO(), filter)
	utils.Checkerr(err)

	for curr.Next(context.TODO()) {
		var product Product
		var prod ViewProduct
		curr.Decode(&product)
		prod.ProductName = product.ProductName
		prod.Price = product.Price
		prod.Quantity = product.Quantity
		products = append(products, prod)
	}

	return products
}
