package service

import (
	"backend/model"
	"context"
)

type Service interface {
	Say(context.Context, string) error
}

type Auth interface {
	SignUp(ctx context.Context, u *model.User) (bool, error)
	Login(ctx context.Context, username string, password string) (string, model.User, error)
}
