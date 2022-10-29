package helper

import (
	dbm "backend/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetConnection() *mongo.Client {
	return dbm.Setup()
}

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	userColl := client.Database("futbol-matches").Collection(collectionName)
	return userColl
}
