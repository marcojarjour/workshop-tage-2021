package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	opsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ops_total",
			Help: "Total operations",
		},
		[]string{"app"},
	)
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of HTTP requests managed",
		},
		[]string{"app", "method"},
	)
	httpRequestsTiming = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_requests_execution_time_milliseconds",
			Help:    "Execution time of HTTP requests in milliseconds",
			Buckets: []float64{1e-10, 1e-8, 1e-6, 1e-4, 1e-2, 0.025, 0.05, 0.075, 0.1, 0.125, 0.25, 0.5, 1, 1.5, 2, 2.5, 5, 7.5, 10, 25, 50, 100, 250, 500, 750, 1000, 2500, 5000, 10000},
		},
		[]string{"app", "method"},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		opsCounter,
		httpRequests,
		httpRequestsTiming,
	)
}

func IncreaseOpsCounter(app string) {
	go opsCounter.WithLabelValues(app).Inc()
}

func IncreaseHttpRequests(app, method string) {
	go httpRequests.WithLabelValues(app, method).Inc()
}

func ObserveHttpRequestsTime(app, method string, timing float64) {
	go httpRequestsTiming.WithLabelValues(app, method).Observe(timing)
}
