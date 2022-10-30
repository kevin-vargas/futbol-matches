package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Ctx = context.TODO()

func Setup() *mongo.Client {
	connectionURI := "mongodb://" + "localhost" + ":" + "27017" + "/"

	credentials := options.Credential{
		Username: "root",
		Password: "q1w2e3r4",
	}

	clientOptions := options.Client().ApplyURI(connectionURI).SetAuth(credentials)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	client.Database("futbol-matches")
	return client
}
