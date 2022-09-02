package middleware

import (
	"backend/pkg/metrics"
	"net/http"
)

func CountRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.RequestCounter.Inc()
		next.ServeHTTP(w, r)
	})
}

/*
func HandleCountRequest(next http.HandlerFunc) http.HandlerFunc {
	h := CountRequest(next)
	return h.ServeHTTP
}
*/
