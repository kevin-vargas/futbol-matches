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
	StartingPlayers  []Player  `bson:"startingPlayers" json:"startingPlayers"`
	SubstitutePlayer []Player  `bson:"substitutePlayer" json:"substitutePlayer"`
}

type Player struct {
	Name  string `bson:"name" json:"name"`
	Phone string `bson:"phone" json:"phone"`
	Email string `bson:"email" json:"email"`
}
