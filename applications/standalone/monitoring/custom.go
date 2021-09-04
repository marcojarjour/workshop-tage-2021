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
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		opsCounter,
	)
}

func IncreaseOpsCounter(app string) {
	go opsCounter.WithLabelValues(app).Inc()
}
