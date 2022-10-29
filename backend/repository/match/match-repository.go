package match

import (
	"backend/model"
	h "backend/repository/helper"
	"backend/service/metrics"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MatchRepository struct{}

func (mr *MatchRepository) GetAllMatches() []model.Match {
	mongoClient := h.GetConnection()
	defer func() {
		mongoClient.Disconnect(h.GetContext())
	}()
	userColl := h.GetCollection(mongoClient, "matches")
	cursor, err := userColl.Find(h.GetContext(), bson.M{}, nil)

	var results = []model.Match{}

	if err == nil {
		for cursor.Next(h.GetContext()) {
			var match model.Match
			err := cursor.Decode(&match)
			if err != nil {
				return []model.Match{}
			}
			results = append(results, match)
		}
	}

	return results
}

func (mr *MatchRepository) GetMatch(id string) model.Match {
	mongoClient := h.GetConnection()
	defer func() {
		mongoClient.Disconnect(h.GetContext())
	}()
	matchColl := h.GetCollection(mongoClient, "matches")
	var match model.Match

	newID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": newID,
	}

	err := matchColl.FindOne(h.GetContext(), condition).Decode(&match)

	if err != nil {
		return model.Match{}
	}
	return match
}

func (mr *MatchRepository) CreateMatch(match model.Match) (string, error) {
	mongoClient := h.GetConnection()
	defer func() {
		mongoClient.Disconnect(h.GetContext())
	}()
	matchesColl := h.GetCollection(mongoClient, "matches")

	match.CreatedAt = time.Now()
	result, err := matchesColl.InsertOne(h.GetContext(), match)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	strAux := strings.Replace(ObjID.String(), "ObjectID(\"", "", 1)
	matchId := strings.Replace(strAux, "\")", "", 1)
	metrics.AddMatch(match)

	return matchId, nil
}

func (mr *MatchRepository) UpdateMatch(match model.Match) error {
	mongoClient := h.GetConnection()
	defer func() {
		mongoClient.Disconnect(h.GetContext())
	}()
	matchColl := h.GetCollection(mongoClient, "matches")

	updateMatch := make(map[string]interface{})

	updateMatch["description"] = match.Description

	updateMatch["finished"] = match.Finished
	updateMatch["date"] = match.Date
	updateMatch["time"] = match.Time
	updateMatch["place"] = match.Place
	updateMatch["format"] = match.Format
	updateMatch["maxPlayers"] = match.MaxPlayers
	updateMatch["startingPlayers"] = match.StartingPlayers
	updateMatch["substitutePlayer"] = match.SubstitutePlayer
	updateMatch["updated_at"] = time.Now()

	updtString := bson.M{
		"$set": updateMatch,
	}

	newId, _ := primitive.ObjectIDFromHex(match.Id)
	filter := bson.M{"_id": newId}

	_, err := matchColl.UpdateOne(h.GetContext(), filter, updtString)
	if err != nil {

		return err
	}
	return nil
}

func (mr *MatchRepository) DeleteMatch(id string) error {
	mongoClient := h.GetConnection()
	defer func() {
		mongoClient.Disconnect(h.GetContext())
	}()
	matchColl := h.GetCollection(mongoClient, "matches")

	newId, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id": newId,
	}

	_, err := matchColl.DeleteOne(h.GetContext(), condition)

	if err != nil {
		return err
	}

	return nil
}

func NewMatchRepository() MatchRepository {
	return MatchRepository{}
}
