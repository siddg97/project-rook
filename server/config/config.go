package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	LogLevel                string
	Env                     string
	Port                    string
	GeminiApiKey            string
	FirebaseCredentialsPath string
}

func InitConfig() (*ServerConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Err(err).Msgf("Error loading .env file")
		panic(err)
	}

	envLogLevel := os.Getenv("LOG_LEVEL")
	envPort := os.Getenv("PORT")
	env := os.Getenv("ENV")
	geminiApiKey := os.Getenv("GEMINI_KEY")

	return &ServerConfig{
		LogLevel:                envLogLevel,
		Env:                     env,
		Port:                    envPort,
		GeminiApiKey:            geminiApiKey,
		FirebaseCredentialsPath: "firebase_credentials.json",
	}, nil
}
