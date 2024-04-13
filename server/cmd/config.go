package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	LogLevel string
	Env      string
	Port     string
}

func InitConfig() (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("Error loading .env file %v", err)
		panic(err)
	}

	envLogLevel := os.Getenv("LOG_LEVEL")
	envPort := os.Getenv("PORT")
	env := os.Getenv("ENV")

	return &ServerConfig{
		LogLevel: envLogLevel,
		Env:      env,
		Port:     envPort,
	}, nil
}
