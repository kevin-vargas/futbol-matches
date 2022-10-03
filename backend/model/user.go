package model

import "time"

// TODO: field validations
type User struct {
	Username  string    `bson:"username" json:"username" validate:"min=5,max=40,required"`
	Password  string    `bson:"password" json:"password" validate:"min=5,max=40,required"`
	Name      string    `bson:"name" json:"name"`
	Phone     string    `bson:"phone" json:"phone"`
	Email     string    `bson:"email" json:"email"`
	CreatedAt time.Time `bson:"created_At" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
