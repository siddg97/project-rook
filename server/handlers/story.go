package handlers

// TO BE REMOVED. POC TO TEST GEMINI INTEGRATION ON THE SERVER

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
)

func GenerateStory(ctx context.Context, client *genai.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if client == nil {
			log.Error().Msg("Client is nil")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Client is null"})
			return
		}
		log.Info().Msgf("Client in generate story %v", client)

		model := client.GenerativeModel("gemini-pro")
		if model == nil {
			log.Error().Msg("Failed to create GenerativeModel")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Model is null"})
			return
		}
		log.Info().Msgf("Model in generate story %v", model)

		prompt_text := genai.Text("Write a story about software engineers suffering in an organization where leadership only cares about how many people goes to the office every week and do not prioritize on moving software to newer architecture.")
		model_response, err := model.GenerateContent(ctx, prompt_text)
		if err != nil {
			log.Fatal().Msgf("Failed to generate content using model %v through Gemini client", model)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"story": model_response})
	}
}
