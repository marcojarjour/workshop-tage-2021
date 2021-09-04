package config

func (c Config) GetEnableMonitoring() bool {
	return c.enableMonitoring
}

func (c Config) GetEnableCustomMetrics() bool {
	return c.enableCustomMetrics
}

func (c Config) GetEnableTracing() bool {
	return c.enableTracing
}

func (c Config) GetTracingTech() string {
	return c.tracingTech
}
