package user

import (
	"backend/model"
	h "backend/repository/helper"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRepository struct{}

func (ur *UserRepository) Create(user model.User) (string, error) {

	existUser := ur.GetByUsername(user.Username)

	if existUser.Username != "" {
		return "", errors.New("The user already exist!")
	}

	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")

	user.CreatedAt = time.Now()
	result, err := userColl.InsertOne(h.GetContext(), user)

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	mongoClient.Disconnect(h.GetContext())

	return ObjID.String(), err
}

func (ur *UserRepository) GetAll() []model.User {
	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")
	cursor, err := userColl.Find(h.GetContext(), bson.M{}, nil)

	var results []model.User

	if err != nil {
		mongoClient.Disconnect(h.GetContext())
		return []model.User{}
	}

	for cursor.Next(h.GetContext()) {
		var user model.User
		err := cursor.Decode(&user)

		if err != nil {
			return results
		}
		results = append(results, user)
	}
	mongoClient.Disconnect(h.GetContext())
	return results
}

func (ur *UserRepository) GetByUsername(username string) model.User {
	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")
	var user model.User

	condition := bson.M{
		"username": username,
	}

	err := userColl.FindOne(h.GetContext(), condition).Decode(&user)

	mongoClient.Disconnect(h.GetContext())
	if err != nil {
		return model.User{}
	}

	return user
}

func (ur *UserRepository) GetByUsernameAndPassword(username string, password string) (model.User, error) {
	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")
	var user model.User

	condition := bson.M{
		"username": username,
		"password": password,
	}

	err := userColl.FindOne(h.GetContext(), condition).Decode(&user)

	mongoClient.Disconnect(h.GetContext())
	if err != nil {
		return model.User{}, errors.New("Incorrect username or password")
	}
	return user, nil
}

func (ur *UserRepository) Update(username string, user model.User) error {
	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")

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

	_, err := userColl.UpdateOne(h.GetContext(), filter, updtString)

	mongoClient.Disconnect(h.GetContext())

	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) Delete(username string) error {

	existUser := ur.GetByUsername(username)

	if existUser.Username == "" {
		return errors.New("User does not exists")
	}

	mongoClient := h.GetConnection()
	userColl := h.GetCollection(mongoClient, "users")

	condition := bson.M{
		"username": username,
	}

	_, err := userColl.DeleteOne(h.GetContext(), condition)

	mongoClient.Disconnect(h.GetContext())

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}
