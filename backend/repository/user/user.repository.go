package user

import (
	dbm "backend/database"
	"backend/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserRepository struct{}

func (ur *UserRepository) Create(user model.User) (string, error) {

	existUser := ur.GetByUsername(user.Username)

	if existUser.Username != "" {
		return "", errors.New("The user already exist!")
	}

	userColl := getCollection("users")
	user.CreatedAt = time.Now()
	result, err := userColl.InsertOne(getContext(), user)
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), err
}

func (ur *UserRepository) GetAll() []model.User {
	userColl := getCollection("users")
	cursor, err := userColl.Find(getContext(), bson.M{}, nil)

	var results []model.User

	if err != nil {
		return []model.User{}
	}

	for cursor.Next(context.TODO()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			return results
		}
		results = append(results, user)
	}
	return results
}

func (ur *UserRepository) GetByUsername(username string) model.User {
	userColl := getCollection("users")
	var user model.User

	condition := bson.M{
		"username": username,
	}

	err := userColl.FindOne(getContext(), condition).Decode(&user)

	if err != nil {
		return model.User{}
	}

	return user
}

func (ur *UserRepository) GetByUsernameAndPassword(username string, password string) (model.User, error) {
	userColl := getCollection("users")
	var user model.User

	condition := bson.M{
		"username": username,
		"password": password,
	}

	err := userColl.FindOne(getContext(), condition).Decode(&user)

	if err != nil {
		return model.User{}, errors.New("Incorrect username or password")
	}

	return user, nil
}

func (ur *UserRepository) Update(username string, user model.User) error {
	userColl := getCollection("users")

	updatedUser := make(map[string]interface{})

	if len(user.Name) > 0 {
		updatedUser["name"] = user.Name
	}

	if len(user.Email) > 0 {
		updatedUser["email"] = user.Email
	}

	if len(user.Phone) > 0 {
		updatedUser["phone"] = user.Phone
	}

	if len(user.Username) > 0 {
		updatedUser["username"] = user.Username
	}

	if len(user.Password) > 0 {
		updatedUser["password"] = user.Password
	}

	updatedUser["updated_at"] = time.Now()

	updtString := bson.M{
		"$set": updatedUser,
	}

	filter := bson.M{"username": bson.M{"$eq": username}}

	_, err := userColl.UpdateOne(getContext(), filter, updtString)

	if err == nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(username string) error {
	userColl := getCollection("users")

	condition := bson.M{
		"username": username,
	}

	_, err := userColl.DeleteOne(getContext(), condition)

	if err != nil {
		return err
	}

	return nil
}

func getContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	return ctx
}

func getCollection(collectionName string) *mongo.Collection {

	dataBaseClient := dbm.Setup()
	userColl := dataBaseClient.Database("futbol-matches").Collection(collectionName)

	return userColl
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}
