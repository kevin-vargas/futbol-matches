package handler

import (
	"backend/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MatchHandler struct {
}

func (matchHandler *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {

	var match model.Match
	json.NewDecoder(r.Body).Decode(&match)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func (matchHandler *MatchHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	someMatch := new(model.Match)
	someMatch.Id = id

	json.NewEncoder(w).Encode(someMatch)
}

func (matchHandler *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	var match model.Match
	json.NewDecoder(r.Body).Decode(&match)
	match.Id = id

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)
}

func (matchHandler *MatchHandler) Delete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("deleted id: " + id)
}

func NewMatchHandler() MatchHandler {
	return MatchHandler{}
}
