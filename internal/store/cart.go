package store

import (
	"context"
	"fmt"

	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartAdd struct {
	ProductName string `json:"product_name"`
	Price       string `json:"price,omitempty"`
	Quantity    int    `json:"quantity"`
}

func UpadateCart(cartproducts []CartAdd) {

	var prod []interface{}
	for num, product := range cartproducts {
		isTrue := checkQuantity(cartproducts[num])
		if isTrue {
			prod = append(prod, product)
		}
	}
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

type tot struct {
	ID    string `bson:"_id"`
	Total int    `bson:"total"`
}
