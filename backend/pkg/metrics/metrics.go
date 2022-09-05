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

var AnnotatedUsers = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "annotated_users",
		Help: "Number of Annotated Users",
	},
)

var CreatedMatches = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "created_matches",
		Help: "Number of Created Matcheds",
	},
)

var RegisteredUsers = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "registered_users",
		Help: "Number of Registered Users",
	},
)

func Register() {
	prometheus.MustRegister(
		RequestCounter,
		AnnotatedUsers,
		CreatedMatches,
		RegisteredUsers,
	)
}

func NewHandler() http.Handler {
	return promhttp.Handler()
}
