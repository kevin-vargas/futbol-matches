package model

import "time"

type Match struct {
	Id               string
	Description      string
	Finished         bool
	Date             time.Time
	Time             string
	Place            string
	Format           int
	MaxPlayers       int
	StartingPlayers  []string
	SubstitutePlayer []string
}
