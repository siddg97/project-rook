package handlers

// TO BE REMOVED. POC TO TEST GEMINI INTEGRATION ON THE SERVER

import (
	"github.com/siddg97/project-rook/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/services"
)

func GenerateStory(geminiService *services.GeminiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info().Msg("Processing request for generating a story")

		promptText := "Write a story about software engineers suffering in an organization where leadership only cares about how many people goes to the office every week and do not prioritize on moving software to newer architecture."

		log.Info().Msg("Sending story prompt to gemini model")
		modelResponse, err := geminiService.PromptGemini(promptText)
		if err != nil {
			log.Err(err).Msg("Failed to generate content through Gemini client")
			c.JSON(http.StatusInternalServerError, &models.ErrorResponse{Message: "Failed to generate content through Gemini"})
			return
		}
		log.Info().Msg("Successfully generated story by prompting gemini")

		c.JSON(http.StatusOK, gin.H{"story": modelResponse})
	}
}
