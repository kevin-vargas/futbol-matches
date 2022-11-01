package handler

import (
	"backend/model"
	"backend/pkg/metrics"
	"backend/service"
	metricsConstants "backend/service/metrics"

	ms "backend/service/match"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MatchHandler struct {
	m            service.MetricStore
	matchService ms.MatchService
}

func (matchHandler *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {

	var match model.Match

	error := json.NewDecoder(r.Body).Decode(&match)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}
	matchCreated, err := matchHandler.matchService.CreateMatch(match)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte("Error creating match"))
	}
	matchHandler.m.Inc(metricsConstants.CREATED_MATCHES)
	metrics.CreatedMatches.Inc()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(matchCreated))
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
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)

}

func (matchHandler *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var changes model.Match
	error := json.NewDecoder(r.Body).Decode(&changes)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	matchUpdated := matchHandler.matchService.UpdateMatch(id, changes)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matchUpdated)

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
	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		return
	}
	added := matchHandler.matchService.AddPlayer(matchId, player)
	if added {
		matchHandler.m.Inc(metricsConstants.ANNOTATED_USERS)
		metrics.AnnotatedUsers.Inc()
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusConflict)
		http.Error(w, "El partido ya completó la cantidad de jugadores", 400)
		return
	}
}

func NewMatchHandler(ms ms.MatchService, m service.MetricStore) MatchHandler {
	return MatchHandler{
		m:            m,
		matchService: ms,
	}
}
