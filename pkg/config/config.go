package config

import (
	"os"
)

type AppConfig struct {
	TelegramAPIKey string
	ConnString     string
}

func New() *AppConfig {
	return &AppConfig{
		TelegramAPIKey: getEnv("TelegramAPIKey", ""),
		ConnString:     getConnString(),
	}
}

/* /pgx/v5 - only
func getConnConfig() *pgx.ConnConfig {
	parseUint, _ := strconv.ParseUint(getEnv("DB_port", "5432"), 10, 32)
	config := pgx.ConnConfig{
		User:     getEnv("DB_user_name", ""),
		Password: getEnv("DB_user_password", ""),
		Database: getEnv("DB_name", ""),
		Host:     getEnv("DB_addr", "localhost"),
		Port:     uint16(parseUint),
	}
	return &config
}*/

func getConnString() string {
	return getEnv("DB_URL", "")
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
