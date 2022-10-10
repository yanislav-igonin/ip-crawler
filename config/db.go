package config

import (
	"os"
)

type DbConfig struct {
	Url string
}

func getValueOrDefaultString(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func getDbConfig() DbConfig {
	url := getValueOrDefaultString(os.Getenv("DATABASE_URL"), "postgresql://localhost/ip_crawler")

	return DbConfig{
		Url: url,
	}
}

var Db DbConfig

func Init() {
	Db = getDbConfig()
}
