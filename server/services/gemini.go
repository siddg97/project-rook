package services

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"google.golang.org/api/option"
)

type GeminiService struct {
	ctx          context.Context
	GeminiClient *genai.Client
	GeminiModel  *genai.GenerativeModel
}

func InitializeGemini(ctx context.Context, geminiKey string) (*GeminiService, error) {
	// Setup Gemini SDK
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiKey))
	if err != nil {
		log.Err(err).Msgf("Error creating Gemini client")
		return nil, err
	}

	model := client.GenerativeModel("gemini-pro")

	geminiService := &GeminiService{
		ctx:          ctx,
		GeminiClient: client,
		GeminiModel:  model,
	}

	log.Info().Msg("Successfully initialized Gemini client")
	return geminiService, nil
}

func (s *GeminiService) PromptGemini(promptText string) (string, error) {
	promptInput := genai.Text(promptText)
	response, err := s.GeminiModel.GenerateContent(s.ctx, promptInput)
	if err != nil {
		return "", err
	}

	var fullResponse string
	for _, candidates := range response.Candidates {
		for _, part := range candidates.Content.Parts {
			textPart := part.(genai.Text)
			fullResponse += string(textPart)
		}
	}

	return fullResponse, nil
}
