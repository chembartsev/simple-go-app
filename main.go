package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	requestCount int
	startTime    = time.Now()
)

func main() {
	// Простой HTTP сервер
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		fmt.Fprintf(w, "<h1>Simple Go App</h1><p>CI/CD is working! 🚀</p><p>Requests: %d</p>", requestCount)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// НОВЫЙ ENDPOINT ДЛЯ PROMETHEUS МЕТРИК
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, `# HELP app_healthy Application health status
# TYPE app_healthy gauge
app_healthy 1

# HELP app_requests_total Total number of requests
# TYPE app_requests_total counter
app_requests_total %d

# HELP app_uptime_seconds Application uptime in seconds
# TYPE app_uptime_seconds gauge
app_uptime_seconds %f
`, requestCount, time.Since(startTime).Seconds())
	})

	fmt.Println("Server starting on :8080...")
	fmt.Println("📊 Metrics available at: http://localhost:8080/metrics")
	http.ListenAndServe(":8080", nil)
}
