package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type Metrics struct {
	httpRequestsCounter *prometheus.CounterVec
}

func NewMetrics() *Metrics {
	m := &Metrics{}
	m.httpRequestsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_counter",
		Help: "number of http requests",
	}, []string{"route"})

	return m
}

func RegisterMetrics(m *Metrics) {
	// create registry.
	registry := prometheus.NewRegistry()

	// register default and custom metrics
	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		m.httpRequestsCounter,
	)

	// Middleware to increment the httpRequestsCounter
	http.Handle("/metrics", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.httpRequestsCounter.WithLabelValues(r.URL.Path).Inc()
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(w, r)
	}))
}
