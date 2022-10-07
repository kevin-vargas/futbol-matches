package handler

import (
	"backend/model"
	"backend/pkg/metrics"
	us "backend/service/user"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserHandler struct {
	userService us.UserService
}

func (userHandler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var user model.User

	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}

	userCreated, error := userHandler.userService.Create(user)

	if error != nil {
		http.Error(w, error.Error(), 400)
		return
	}
	metrics.CreatedUsers.Inc()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(userCreated))

}

func (userHandler *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := userHandler.userService.GetAll()
	if len(response) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func (userHandler *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	w.Header().Set("Content-Type", "application/json")

	user := userHandler.userService.GetByUsername(username)
	if user.Username == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (userHandler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	var user model.User
	error := json.NewDecoder(r.Body).Decode(&user)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := userHandler.userService.Update(username, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (userHandler *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	w.Header().Set("Content-Type", "application/json")

	err := userHandler.userService.Delete(username)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("deleted user: " + username)
	}

}

func NewUserHandler(us us.UserService) UserHandler {
	return UserHandler{
		userService: us,
	}
}
