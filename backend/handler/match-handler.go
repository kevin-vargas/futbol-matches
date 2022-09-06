package handler

import (
	"backend/model"
	"backend/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MatchHandler struct {
	matchService service.MatchService
}

func (matchHandler *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var match model.Match

	error := json.NewDecoder(r.Body).Decode(&match)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	} else {
		matchCreated := matchHandler.matchService.CreateMatch(match)
		if matchCreated.Id != "" {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(matchCreated.Id)
		} else {
			w.WriteHeader(http.StatusNotModified)
		}
	}
}

func (matchHandler *MatchHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(matchHandler.matchService.GetAllMatches())
}

func (matchHandler *MatchHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	match := matchHandler.matchService.GetMatch(id)
	if match.Id == "" {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(match)
	}
}

func (matchHandler *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var changes model.Match
	error := json.NewDecoder(r.Body).Decode(&changes)

	w.Header().Set("Content-Type", "application/json")
	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	} else {
		matchUpdated := matchHandler.matchService.UpdateMatch(id, changes)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(matchUpdated)
	}
}

func (matchHandler *MatchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	matchHandler.matchService.DeleteMatch(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("deleted id: " + id)
}

func (matchHandler *MatchHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	matchId := chi.URLParam(r, "id")
	var player model.Player
	error := json.NewDecoder(r.Body).Decode(&player)

	w.Header().Set("Content-Type", "application/json")
	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	} else {
		added := matchHandler.matchService.AddPlayer(matchId, player)
		if added {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Player added")
		} else {
			http.Error(w, "El partido ya completó la cantidad de jugadores", 400)
			return
		}
	}
}

func NewMatchHandler(ms service.MatchService) MatchHandler {
	return MatchHandler{
		matchService: ms,
	}
}
