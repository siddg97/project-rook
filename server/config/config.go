package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	LogLevel     string
	Env          string
	Port         string
	GeminiApiKey string
}

func InitConfig(ctx context.Context) (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msgf("Error loading .env file")
		log.Warn().Msg("No .env file found, will try to source configuration from runtime environment variables")
	}

	envLogLevel := getDefaultEnvVar("LOG_LEVEL", "debug")
	envPort := getDefaultEnvVar("PORT", "3000")
	envEnv := getDefaultEnvVar("ENV", "local")
	envGeminiKey, err := getEnvVar("GEMINI_KEY")
	if err != nil {
		log.Err(err).Msgf("Invalid configuration")
		panic(err)
	}

	return &ServerConfig{
		LogLevel:     envLogLevel,
		Env:          envEnv,
		Port:         envPort,
		GeminiApiKey: envGeminiKey,
	}, nil
}

func getEnvVar(name string) (string, error) {
	envValue := os.Getenv(name)
	if envValue == "" {
		invalidConfigError := errors.New(fmt.Sprintf("Required configuration value for '%s' not found", name))
		return "", invalidConfigError
	}

	return envValue, nil
}

func getDefaultEnvVar(name string, defaultValue string) string {
	envValue := os.Getenv(name)
	if envValue == "" {
		log.Info().Msgf("No %s env var found, defaulting to '%s'", name, defaultValue)
		return defaultValue
	}
	return envValue
}
