package config

import (
	"os"
	"strconv"
)

type Config struct {
	ProjectName      string
	StdOutLogTimeout int
	Metrics          MetricsConfig
}

type MetricsConfig struct {
	Method          string
	Route           string
	Port            int
	Component       string
	Source          string
	Namespace       string
	CollectInterval int
}

func GetConfig() *Config {
	return &Config{
		ProjectName:      getEnv("PROJECT_NAME", "go-test-app"),
		StdOutLogTimeout: getEnvInt("CONSOLE_LOG_TIMEOUT", 5000),
		Metrics: MetricsConfig{
			Method:          getEnv("METRICS_METHOD", "GET"),
			Route:           getEnv("METRICS_ROUTE", "/metrics"),
			Port:            getEnvInt("METRICS_PORT", 8080),
			Component:       getEnv("METRICS_COMPONENT", "applications"),
			Source:          getEnv("METRICS_SOURCE", "go-test-app"),
			Namespace:       getEnv("METRICS_NAMESPACE", ""),
			CollectInterval: getEnvInt("METRICS_COLLECT_INTERVAL", 10),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return fallback
}
