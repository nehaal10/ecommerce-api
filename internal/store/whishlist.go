package store

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Wishlist_add(wish []Wishlist, id string) []string {
	//check
	var result []string
	var update primitive.M
	var user CustomerDatabase
	var addwish []CartAdd
	var w []interface{}
	filter := bson.M{
		"user_id": id,
	}
	db.Cutomer.FindOne(context.TODO(), filter).Decode(&user)
	cartprods := user.Cart
	for _, val := range wish {
		var wishlist CartAdd
		for _, prod := range cartprods {
			if val.Product_name == prod.ProductName {
				wishlist.ProductName = val.Product_name
				wishlist.Price = prod.Price
				wishlist.Quantity = prod.Quantity
				addwish = append(addwish, wishlist)
				w = append(w, wishlist)
				str := fmt.Sprintf("%s added to wishlist", val.Product_name)
				result = append(result, str)
			}
		}
	}

	if len(user.ProductWishList) == 0 {
		update = bson.M{
			"$set": bson.M{
				"wishlist": w,
			},
		}
	} else {
		update = bson.M{
			"$push": bson.M{
				"wishlist": bson.M{
					"$each": w,
				},
			},
		}
	}

	db.Cutomer.UpdateOne(context.TODO(), filter, update)
	_ = DeleteProduct(addwish, id)
	return result
}

func ViewWishlist(id string) []CartAdd {
	var user CustomerDatabase
	filter := bson.M{
		"user_id": id,
	}

	db.Cutomer.FindOne(context.TODO(), filter).Decode(&user)

	return user.ProductWishList
}

func DeleteWishlist(wishProd []Wishlist, id string) []string {
	filter := bson.M{"user_id": id}
	var response []string

	for _, prod := range wishProd {
		if checkwishproducts(prod.Product_name, id) {
			update := bson.M{
				"$pull": bson.M{
					"wishlist": bson.M{
						"productname": prod.Product_name,
					},
				},
			}

			db.Cutomer.UpdateOne(context.TODO(), filter, update)
			str := fmt.Sprintf("%s removed from wish list", prod.Product_name)
			response = append(response, str)
		} else {
			str := fmt.Sprintf("%s not there in the wishlist", prod.Product_name)
			response = append(response, str)
		}
	}
	return response
}

func checkwishproducts(productname string, id string) bool {
	var u CustomerDatabase
	filter := bson.M{
		"user_id": id,
	}
	res := db.Cutomer.FindOne(context.TODO(), filter)
	res.Decode(&u)
	for _, prods := range u.ProductWishList {
		if productname == prods.ProductName {
			return true
		}
	}
	return false
}
