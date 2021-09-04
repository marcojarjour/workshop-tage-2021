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
	succConsumedMsgCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "success_consumed_messages_total",
			Help: "Total SUCCESSFULLY consumed messages",
		},
		[]string{"app", "topic"},
	)
	failConsumedMsgCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fail_consumed_messages_total",
			Help: "Total FAILED consumed messages",
		},
		[]string{"app", "topic"},
	)
	consumerErrCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "consumer_errors_total",
			Help: "Total consumer errors",
		},
		[]string{"app"},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		opsCounter,
		succConsumedMsgCounter,
		failConsumedMsgCounter,
		consumerErrCounter,
	)
}

func IncreaseOpsCounter(app string) {
	go opsCounter.WithLabelValues(app).Inc()
}

func IncreaseSuccConsumedMsgCounter(app, topic string) {
	go succConsumedMsgCounter.WithLabelValues(app, topic).Inc()
}

func IncreaseFailConsumedMsgCounter(app, topic string) {
	go failConsumedMsgCounter.WithLabelValues(app, topic).Inc()
}

func IncreaseConsumerErrCounter(app string) {
	go consumerErrCounter.WithLabelValues(app).Inc()
}
