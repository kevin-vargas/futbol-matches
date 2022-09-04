package handler

import (
	"backend/model"
	"backend/pkg/metrics"
	"backend/service"
	"net/http"
)

type Auth struct {
	service service.Auth
}

func (a *Auth) SingUp(w http.ResponseWriter, r *http.Request) {
	user, err := decode[model.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ok, err := a.service.SignUp(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !ok {
		w.WriteHeader(http.StatusConflict)
		return
	}
	metrics.RegisteredUsers.Inc()
	metrics.CreatedMatches.Inc()
	metrics.AnnotatedUsers.Inc()
	w.WriteHeader(http.StatusCreated)
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	u, err := decode[model.User](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := a.service.Login(r.Context(), u.Username, u.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid username or password"))
		return
	}
	w.Header().Add("Authorization", "Bearer "+token)
	w.WriteHeader(http.StatusNoContent)
}

func NewAuth(service service.Auth) Auth {
	return Auth{
		service: service,
	}
}
