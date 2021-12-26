package db

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection
var CTX = context.TODO()

func Connect(host string, port int) *mongo.Client {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s/", host, strconv.Itoa(port)))
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(CTX, nil)
	if err != nil {
		log.Fatal(err)
	}

	Collection = client.Database("pellets").Collection("prices")

	return client
}
