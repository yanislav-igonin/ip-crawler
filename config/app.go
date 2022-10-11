package config

import (
	"os"
	"strconv"
)

type AppConfig struct {
	RequestIntervalSeconds int
}

func getValueOrDefaultInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return res
}

func getAppConfig() AppConfig {
	RequestIntervalSeconds := getValueOrDefaultInt(os.Getenv("REQUEST_INTERVAL_SECONDS"), 1)

	return AppConfig{
		RequestIntervalSeconds,
	}
}

var App AppConfig

func initAppConfig() {
	App = getAppConfig()
}
