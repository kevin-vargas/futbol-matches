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
	w.Write([]byte("Bearer " + token))
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	u, err := decode[model.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := a.service.Login(u.Username, u.Password)

	w.Header().Set("Content-Type", "application/json")

	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid username or password"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func NewAuth(service user.UserService) Auth {
	return Auth{
		service: service,
	}
}
