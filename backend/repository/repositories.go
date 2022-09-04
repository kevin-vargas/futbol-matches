package repository

import (
	"backend/model"
)

type User = Repository[model.User]

var NewUser = NewInMemory[model.User]
