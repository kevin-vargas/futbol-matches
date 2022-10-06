package helper

import (
	dbm "backend/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx
}

func GetCollection(collectionName string) *mongo.Collection {

	dataBaseClient := dbm.Setup()
	userColl := dataBaseClient.Database("futbol-matches").Collection(collectionName)

	return userColl
}
