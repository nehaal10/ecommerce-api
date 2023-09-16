package store

import (
	"context"

	"github.com/nehaal10/ecommerce-api/internal/conf"
	"github.com/nehaal10/ecommerce-api/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Collection struct {
	NewUser *mongo.Collection
	Cutomer *mongo.Collection
	Product *mongo.Collection
}

var db Collection

func SetupDBConeection(cfg conf.Config) {
	clientoption := options.Client().ApplyURI(cfg.DbHost)
	client, err := mongo.Connect(context.TODO(), clientoption)
	utils.Checkerr(err)
	dbLogin := client.Database("ecommerce").Collection("auth")
	dbCustomer := client.Database("ecommerce").Collection("user")
	dbProduct := client.Database("ecommerce").Collection("product")
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	db = Collection{
		NewUser: dbLogin,
		Cutomer: dbCustomer,
		Product: dbProduct,
	}

}

func GetdbConnection() Collection {
	return db
}
