package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/models"
	"github.com/siddg97/project-rook/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func GetUserSummary(firebaseService *services.FirebaseService, geminiService *services.GeminiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Source userId from path param
		userId := c.Param("userId")
		log.Info().Msgf("Processing resume summary request for user: %s", userId)

		// Fetch prompt history for user
		promptHistory, err := firebaseService.GetResumePromptHistory(userId)
		if err != nil {
			log.Err(err).Msgf("Could not fetch resume prompt history for user: %s", userId)

			if status.Code(err) == codes.NotFound {
				log.Info().Msgf("Did not find any resume data for user: %s", userId)
				c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "No resume data found"})
				return
			}
			c.JSON(http.StatusNotFound, models.ErrorResponse{Message: "Something failed while getting data"})
			return
		}

		// Get gemini to summarize the resume
		summarizePromptText := "summarize the resume provided to you earlier. In your response, only provide the markdown version fo the resume summary"
		summary, err := geminiService.PromptGeminiWithHistory(promptHistory, summarizePromptText)
		if err != nil {
			log.Err(err).Msgf("Error while generating resume summary for user: %s", userId)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error generating summary"})
			return
		}

		// Save the summary prompt and gemini response to prompt history in db
		err = firebaseService.StoreToPromptHistory(userId, summarizePromptText, "user")
		if err != nil {
			log.Err(err).Msg("Could not store summarize resume prompt for gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing summarize resume prompt in db"})
			return
		}
		err = firebaseService.StoreToPromptHistory(userId, summary, "model")
		if err != nil {
			log.Err(err).Msg("Could not store summarize resume response from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing summarize resume prompt response in db"})
			return
		}

		c.JSON(http.StatusOK, summary)

	}

}
