package config

import "os"

type TelegramApiConfig struct {
	APIKey string
}

func New() *TelegramApiConfig {
	return &TelegramApiConfig{
		APIKey: getEnv("TelegramAPIKey", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
