package config

import "os"

type Config struct {
	DBPath    string
	JWTSecret string
	Port      string
}

func LoadConfig() *Config {
	return &Config{
		DBPath:    getEnv("DB_PATH", "golearn.db"),
		JWTSecret: getEnv("JWT_SECRET", "super-secret-key-golearn-2024"),
		Port:      getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
