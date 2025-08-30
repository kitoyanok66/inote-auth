package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret   string
	JWTDuration time.Duration

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	jwtDurationStr := os.Getenv("JWT_DURATION")
	if jwtDurationStr == "" {
		jwtDurationStr = "15m"
	}
	jwtDuration, err := time.ParseDuration(jwtDurationStr)
	if err != nil {
		log.Fatalf("Invalid JWT_DURATION: %v", err)
	}

	return &Config{
		JWTSecret:   jwtSecret,
		JWTDuration: jwtDuration,
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5433"),
		DBUser:      getEnv("DB_USER", "auth_user"),
		DBPassword:  getEnv("DB_PASSWORD", "auth_pass"),
		DBName:      getEnv("DB_NAME", "auth"),
		DBSSLMode:   getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
