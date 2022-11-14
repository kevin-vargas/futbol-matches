package handler

import (
	"backend/model"
	"backend/pkg/metrics"
	"backend/service"
	metricsConstants "backend/service/metrics"
	us "backend/service/user"

	ms "backend/service/match"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MatchHandler struct {
	m            service.MetricStore
	matchService ms.MatchService
	userService  us.UserService
}

func (matchHandler *MatchHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var match model.Match

	error := json.NewDecoder(r.Body).Decode(&match)

	if error != nil {
		http.Error(w, "Error en los datos recibidos "+error.Error(), 400)
		json.NewEncoder(w).Encode(map[string]string{"error": error.Error()})
		return
	}
	matchCreated, err := matchHandler.matchService.CreateMatch(match)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(map[string]string{"error": error.Error()})
	}

	user := matchHandler.userService.GetByUsername(match.Owner)
	player := model.Player{Name: user.Username, Phone: user.Phone, Email: user.Email}

	matchHandler.matchService.AddPlayer(matchCreated, player)
	matchHandler.m.Inc(metricsConstants.ANNOTATED_USERS)
	matchHandler.m.Inc(metricsConstants.CREATED_MATCHES)
	metrics.CreatedMatches.Inc()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"matchId": matchCreated})
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
		json.NewEncoder(w).Encode(map[string]string{"error": "match not found"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(match)

}

func (matchHandler *MatchHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	var changes model.Match
	error := json.NewDecoder(r.Body).Decode(&changes)

	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": error.Error()})
		return
	}

	matchUpdated := matchHandler.matchService.UpdateMatch(id, changes)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matchUpdated)

}

func (matchHandler *MatchHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	matchHandler.matchService.DeleteMatch(id)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"deleted": id})
}

func (matchHandler *MatchHandler) AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	matchId := chi.URLParam(r, "id")
	var player model.Player
	error := json.NewDecoder(r.Body).Decode(&player)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": error.Error()})
		return
	}

	added, err := matchHandler.matchService.AddPlayer(matchId, player)
	if added {
		matchHandler.m.Inc(metricsConstants.ANNOTATED_USERS)
		metrics.AnnotatedUsers.Inc()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(player)
		return
	} else {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
}

func NewMatchHandler(ms ms.MatchService, m service.MetricStore, us us.UserService) MatchHandler {
	return MatchHandler{
		m:            m,
		matchService: ms,
		userService:  us,
	}
}
