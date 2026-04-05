package configs

import (
	"os"
)

type Config struct {
	AppName string
	AppEnv  string
	AppPort string
	JWTSecret string
	TokenHours string
}

func LoadConfig() Config {
	return Config{
		AppName: os.Getenv("APP_NAME", "diotron"),
		AppEnv:  os.Getenv("APP_ENV", "development"),
		AppPort: os.Getenv("APP_PORT", "8081"),
		JWTSecret: os.Getenv("JWT_SECRET", "super-secret-diotron-key"),
		TokenHours: os.Getenv("TOKEN_HOURS", "24"),
	}
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}