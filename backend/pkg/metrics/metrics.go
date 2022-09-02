package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var RequestCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "request_count",
		Help: "Number of request handled",
	},
)

func Register() {
	prometheus.MustRegister(RequestCounter)
}

func NewHandler() http.Handler {
	return promhttp.Handler()
}
