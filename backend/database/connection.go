package database

import (
	"backend/config"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Ctx = context.TODO()

func Setup() *mongo.Client {
	dbConfig := config.New().DB

	connectionURI := dbConfig.ConnectionURI

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().ApplyURI(connectionURI).SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	client.Database(dbConfig.Database)
	return client
}
