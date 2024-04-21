package handlers

// TO BE REMOVED. POC TO TEST GEMINI INTEGRATION ON THE SERVER

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/services"
)

func GenerateStory(c *gin.Context) {
	ctx := context.Background()

	geminiClient := services.GetGeminiClient()

	model := geminiClient.GenerativeModel("gemini-pro")

	promptText := genai.Text("Write a story about software engineers suffering in an organization where leadership only cares about how many people goes to the office every week and do not prioritize on moving software to newer architecture.")
	modelResponse, err := model.GenerateContent(ctx, promptText)
	if err != nil {
		log.Fatal().Msgf("Failed to generate content using model %v through Gemini client", model)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"story": modelResponse})
}
