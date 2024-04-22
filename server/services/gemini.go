package services

import (
	"context"
	"fmt"
	"github.com/siddg97/project-rook/models"

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

	model := client.GenerativeModel("gemini-1.0-pro")

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

	return responseToString(response), nil
}

func (s *GeminiService) PromptGeminiWithHistory(history []models.PromptHistoryDocument, promptText string) (string, error) {
	// Start chat session
	chatSession := s.GeminiModel.StartChat()

	// Set chat session history
	var promptHistoryContent []*genai.Content
	for _, promptHistory := range history {
		promptHistoryContent = append(promptHistoryContent, &genai.Content{
			Parts: []genai.Part{
				genai.Text(promptHistory.Text),
			},
			Role: promptHistory.Role,
		})
	}
	chatSession.History = promptHistoryContent

	for i, c := range chatSession.History {
		log.Info().Msgf("%d: %+v", i, c)
	}

	// Send prompt to gemini with chat history
	response, err := chatSession.SendMessage(s.ctx, genai.Text(promptText))
	if err != nil {
		fmt.Printf("%v", err)
		return "", err
	}

	return responseToString(response), nil
}

func responseToString(response *genai.GenerateContentResponse) string {
	var fullResponse string
	for _, candidates := range response.Candidates {
		for _, part := range candidates.Content.Parts {
			textPart := part.(genai.Text)
			fullResponse += string(textPart)
		}
	}
	return fullResponse
}
