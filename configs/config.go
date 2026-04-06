package configs

import (
	"os"
)

type Config struct {
	AppName    string
	AppEnv     string
	AppPort    string
	JWTSecret  string
	TokenHours string
}

func LoadConfig() Config {
	return Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppPort:    os.Getenv("APP_PORT"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		TokenHours: os.Getenv("TOKEN_HOURS"),
	}
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
