package handler

import (
	"backend/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Metric struct {
	m service.Metric
}

func (h *Metric) GetLastMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	metric := chi.URLParam(r, "metric")
	query := r.URL.Query().Get("query")
	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "send query param"})
		return
	}
	result, err := h.m.LastMetrics(r.Context(), metric, query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(map[string]string{"error": "Error on query metric"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"count": strconv.Itoa(result)})
}

func NewMetric(m service.Metric) Metric {
	return Metric{
		m: m,
	}
}
