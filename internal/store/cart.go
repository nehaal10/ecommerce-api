package store

import (
	"context"
	"fmt"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartAdd struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price,omitempty"`
	Quantity    int    `json:"quantity"`
}

func UpadateCart(cartproducts []CartAdd, id string) {
	var cus CustomerDatabase
	var update primitive.M
	var prod []interface{}
	for num, product := range cartproducts {
		ok, price := checkProductInventory(cartproducts[num].ProductName)
		if !ok {
			isTrue := checkQuantity(cartproducts[num])
			if isTrue {
				product.Price = price
				prod = append(prod, product)
			}
		}
	}

	filter := bson.M{
		"user_id": id,
	}

	db.Cutomer.FindOne(context.TODO(), filter).Decode(&cus)

	if cus.Cart == nil {
		update = bson.M{
			"$set": bson.M{
				"cart": prod,
			},
		}
	} else {
		update = bson.M{
			"$push": bson.M{
				"cart": bson.M{
					"$each": prod,
				},
			},
		}
	}

	db.Cutomer.UpdateOne(context.TODO(), filter, update)
}

func checkQuantity(product CartAdd) bool {
	return aggregateQuantity(product.ProductName, product.Quantity)
}

func aggregateQuantity(prod string, q int) bool {
	var t []tot
	matchstage := bson.D{
		{Key: "$match", Value: bson.D{
			{
				Key: "product_name", Value: prod,
			},
		},
		},
	}

	groupstage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: prod},
			{Key: "total", Value: bson.D{
				{Key: "$sum", Value: "$quantity"},
			}},
		}},
	}
	curr, err := db.Product.Aggregate(context.Background(), mongo.Pipeline{matchstage, groupstage})
	utils.Checkerr(err)
	if err := curr.All(context.TODO(), &t); err != nil {
		fmt.Println("hello")
		panic(err)
	}
	return t[0].Total >= q
}

func checkProductInventory(prodName string) (bool, int) {
	var p []Product
	filter := bson.M{
		"product_name": prodName,
	}

	curr, err := db.Product.Find(context.TODO(), filter)
	utils.Checkerr(err)
	curr.All(context.TODO(), &p)
	price := p[0].Price
	return p == nil, price
}

type tot struct {
	ID    string `bson:"_id"`
	Total int    `bson:"total"`
}

func ViewCart(id string) []CartAdd {
	var user CustomerDatabase
	filter := bson.M{
		"user_id": id,
	}
	result := db.Cutomer.FindOne(context.TODO(), filter)
	result.Decode(&user)
	cart := user.Cart
	return cart
}

func checkcartproducts(productname string, id string) bool {
	var u CustomerDatabase
	filter := bson.M{
		"user_id": id,
	}
	res := db.Cutomer.FindOne(context.TODO(), filter)
	res.Decode(&u)
	for _, prods := range u.Cart {
		if productname == prods.ProductName {
			return true
		}
	}
	return false
}

func DeleteProduct(cartproducts []CartAdd, id string) []string {
	filter := bson.M{"user_id": id}
	var response []string
	for _, val := range cartproducts {
		ok := checkcartproducts(val.ProductName, id)
		if ok {
			update := bson.M{
				"$pull": bson.M{
					"cart": bson.M{
						"productname": val.ProductName,
					},
				},
			}
			db.Cutomer.UpdateOne(context.TODO(), filter, update)
			message := fmt.Sprintf("%s deleted successfully", val.ProductName)
			response = append(response, message)
		} else {
			str := fmt.Sprintf("%s is not there in the cart", val.ProductName)
			response = append(response, str)
		}
	}

	return response
}

//take care of json returns for adding product to cart
