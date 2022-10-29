package mongodb

import (
	"backend/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Ctx = context.TODO()

func Setup() *mongo.Client {
	dbConfig := config.New().DB
	connectionURI := "mongodb://" + dbConfig.Host + ":" + dbConfig.Port + "/"

	credentials := options.Credential{
		Username: dbConfig.User,
		Password: dbConfig.Pass,
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
