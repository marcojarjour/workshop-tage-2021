package config

type Config struct {
	enableMonitoring    bool
	enableCustomMetrics bool
	enableTracing       bool
	tracingTech         string
}
