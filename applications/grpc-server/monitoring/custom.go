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
	greetingsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "greetings_total",
			Help: "Total greetings",
		},
		[]string{"app", "name"},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		opsCounter,
		greetingsCounter,
	)
}

func IncreaseOpsCounter(app string) {
	go opsCounter.WithLabelValues(app).Inc()
}

func IncreaseGreetingsCounter(app, name string) {
	go greetingsCounter.WithLabelValues(app, name).Inc()
}
