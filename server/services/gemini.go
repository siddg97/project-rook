package services

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

func InitializeGemini(ctx context.Context, geminiKey string) (*genai.Client, error) {
	// Setup Gemini SDK
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiKey))
	if err != nil {
		log.Fatal().Msgf("Error creating Gemini client: %v", err)
		return nil, err
	}

	log.Info().Msg("Successfully initialized Gemini client")
	return client, nil
}
