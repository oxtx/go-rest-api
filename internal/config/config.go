package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	AppName      string
	Env          string
	HTTPPort     string
	DBURL        string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func Load() *Config {
	return &Config{
		AppName:      getEnv("APP_NAME", "go-rest-api"),
		Env:          getEnv("APP_ENV", "development"),
		HTTPPort:     getEnv("HTTP_PORT", "8080"),
		DBURL:        getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/awesome?sslmode=disable"),
		ReadTimeout:  getEnvDuration("HTTP_READ_TIMEOUT", 5),
		WriteTimeout: getEnvDuration("HTTP_WRITE_TIMEOUT", 10),
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvDuration(key string, defSeconds int) time.Duration {
	if v := os.Getenv(key); v != "" {
		d, err := time.ParseDuration(v)
		if err == nil {
			return d
		}
		log.Printf("invalid duration for %s using default", key)
	}
	return time.Duration(defSeconds) * time.Second
}
