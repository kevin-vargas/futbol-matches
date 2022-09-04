package service

import (
	"backend/model"
	"backend/repository"
	"context"
	"errors"
)

type Encrypt interface {
	Generate(string) (string, error)
	Compare(string, string) (bool, error)
}

type JWT interface {
	Generate(string) (string, error)
}

type auth struct {
	r repository.User
	e Encrypt
	j JWT
}

func (a *auth) SignUp(ctx context.Context, u *model.User) (bool, error) {
	user, err := a.r.FindOne(u.Username)
	if err != nil {
		return false, errors.New("error on repository")
	}
	if user != nil {
		return false, nil
	}
	p, err := a.e.Generate(u.Password)
	if err != nil {
		return false, err
	}
	u.Password = string(p)
	err = a.r.Save(u.Username, u)
	if err != nil {
		return false, errors.New("error on user save")
	}
	return true, nil
}

func (a *auth) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := a.r.FindOne(username)
	if err != nil {
		return "", errors.New("error on repository")
	}
	if user != nil {
		if ok, err := a.e.Compare(user.Password, password); err == nil && ok {
			return a.j.Generate(user.Username)
		}
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func NewAuth(r repository.User, e Encrypt, j JWT) Auth {
	return &auth{
		r: r,
		e: e,
		j: j,
	}
}
