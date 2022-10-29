package handler

import (
	"backend/model"
	"backend/pkg/metrics"
	"backend/service/user"
	"encoding/json"
	"net/http"
)

type Auth struct {
	service user.UserService
}

func (a *Auth) SingUp(w http.ResponseWriter, r *http.Request) {
	user, err := decode[model.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := a.service.Create(*user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token == "" {
		w.WriteHeader(http.StatusConflict)
		return
	}
	metrics.RegisteredUsers.Inc()
	w.Header().Add("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusCreated)
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	u, err := decode[model.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := a.service.Login(u.Username, u.Password)

	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username or password"))
		return
	}

	w.Header().Add("Authorization", "Bearer "+token)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(token)
}

func NewAuth(service user.UserService) Auth {
	return Auth{
		service: service,
	}
}
