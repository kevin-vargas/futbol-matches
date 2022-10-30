package handler

import (
	ms "backend/service/metrics"
	"encoding/json"
	"net/http"
)

type MetricsHandler struct {
	metricsService ms.MetricsService
}

func (metricsHandler *MetricsHandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches := metricsHandler.metricsService.GetLastCreatedMatches()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(len(matches))

}

func (metricsHandler *MetricsHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players := metricsHandler.metricsService.GetLastJoinedPlayers()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(len(players))
}

func NewMetricsHandler(ms ms.MetricsService) MetricsHandler {
	return MetricsHandler{
		metricsService: ms,
	}
}
