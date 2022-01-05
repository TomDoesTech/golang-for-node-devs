package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var mongoOnce sync.Once

var clientInstanceError error

type Collection string

const (
	ProductsCollection Collection = "products"
)

const (
	url      = "mongodb://localhost:27017"
	Database = "products-api"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(url)

		client, err := mongo.Connect(context.TODO(), clientOptions)

		clientInstance = client

		clientInstanceError = err
	})

	return clientInstance, clientInstanceError
}
