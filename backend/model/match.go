package model

import "time"

type Match struct {
	Id               string    `bson:"_id" json:"id"`
	Description      string    `bson:"description" json:"description"`
	Finished         bool      `bson:"finished" json:"finished"`
	Date             time.Time `bson:"date" json:"date"`
	Time             string    `bson:"time" json:"time"`
	Place            string    `bson:"place" json:"place"`
	Format           int       `bson:"format" json:"format"`
	MaxPlayers       int       `bson:"maxPlayers" json:"maxPlayers"`
	StartingPlayers  []string  `bson:"startingPlayers" json:"startingPlayers"`
	SubstitutePlayer []string  `bson:"substitutePlayer" json:"substitutePlayer"`
}
