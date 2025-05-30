package config

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DbHost     string `env:"DB_HOST,required"`
	DbName     string `env:"DB_NAME,required"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbSSLMode  string `env:"DB_SSLMODE,required"`
}

func NewEnvConfig() *EnvConfig {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Unable to load .env file: %v", err)
	}
	config := &EnvConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from the .env: %e", err)
	}
	return config
}