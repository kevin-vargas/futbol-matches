package handler

import (
	"net/http"

	"backend/service"

	"github.com/go-chi/chi"
)

type Hello struct {
	service service.Service
}

func (handler *Hello) Handle(w http.ResponseWriter, r *http.Request) {

	word := chi.URLParam(r, "word")
	if word == "" {
		http.Error(w, "word path param not provided", http.StatusBadRequest)
	}
	err := handler.service.Say(r.Context(), word)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func New(s service.Service) Hello {
	return Hello{
		service: s,
	}
}
