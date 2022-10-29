package service

import (
	"backend/model"
	"backend/repository"
	"context"
	"encoding/json"
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

func (a *auth) Login(ctx context.Context, username string, password string) (string, model.User, error) {
	user, err := a.r.FindOne(username)
	bytes, err := json.Marshal(user)

	if err != nil {
		return "", model.User{}, errors.New("error on repository")
	}
	if user != nil {
		if ok, err := a.e.Compare(user.Password, password); err == nil && ok {

			userModel := model.User{
				Username: user.Username,
				Password: user.Username,
				Name:     user.Name,
				Lastname: user.Lastname,
				Phone:    user.Phone,
				Email:    user.Email,
			}

			str, e := a.j.Generate(string(bytes))
			return str, userModel, e
		}
		if err != nil {
			return "", model.User{}, err
		}
	}
	return "", model.User{}, nil
}

func NewAuth(r repository.User, e Encrypt, j JWT) Auth {
	return &auth{
		r: r,
		e: e,
		j: j,
	}
}
