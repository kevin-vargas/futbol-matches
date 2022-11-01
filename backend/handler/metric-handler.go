package handler

import (
	"backend/service"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Metric struct {
	m service.Metric
}

func (h *Metric) GetLastMetrics(w http.ResponseWriter, r *http.Request) {
	metric := chi.URLParam(r, "metric")
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "send query param", 400)
		return
	}
	result, err := h.m.LastMetrics(r.Context(), metric, query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		w.Write([]byte("Error on query metric"))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(result)))
}

func NewMetric(m service.Metric) Metric {
	return Metric{
		m: m,
	}
}
