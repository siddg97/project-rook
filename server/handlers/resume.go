package handlers

import (
	"fmt"
	"github.com/siddg97/project-rook/models"
	"github.com/siddg97/project-rook/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func CreateResume(visionService *services.VisionService, firebaseService *services.FirebaseService, geminiService *services.GeminiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Source userId from path param
		userId := c.Param("userId")
		log.Info().Msgf("Processing resume creation request for user: %s", userId)

		// Source PDF file from request
		requestFile, err := c.FormFile("file")
		if err != nil {
			log.Error().Msgf("Could not find `file` attached from incoming request: %v", err)
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "No file found in request"})
			return
		}

		openedFile, _ := requestFile.Open()
		filename := requestFile.Filename
		log.Info().Msgf("Found file in %v request", filename)

		pdfContent, err := io.ReadAll(openedFile)
		if err != nil {
			log.Error().Msgf("Error to read file with name %s: %v", filename, err)
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Failed to read file"})
			return
		}

		// Extract text from PDF file via GCP Vision APIs
		extractedResumeText, err := visionService.ExtractTextFromPdf(pdfContent)
		if err != nil {
			log.Error().Msgf("Encountered error when trying to extract text from PDF file: %v", err)
			c.JSON(http.StatusInternalServerError, &models.ErrorResponse{Message: "Failed to extract text from resume"})
			return
		}
		log.Info().Msgf("Successfully extracted %v characters of text for attached resume for userId: %v", len(extractedResumeText), userId)

		// Store resume text to prompt history for user
		err = firebaseService.StoreNewResume(userId, extractedResumeText)
		if err != nil {
			log.Err(err).Msg("Could not store resume to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume in db"})
			return
		}

		// Setup context with initial prompt to gemini
		initalContextPrompt := fmt.Sprintf("%s\n\n data to update relevant sections of this resume will be provided", extractedResumeText)
		geminiResponse, err := geminiService.PromptGemini(initalContextPrompt)
		if err != nil {
			log.Err(err).Msg("Failed to save context for resume via gemini prompt")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error prompting gemini for initial resume context"})
			return
		}

		// Store gemini prompt and response to prompt history in db
		err = firebaseService.StoreToPromptHistory(userId, initalContextPrompt, "user")
		if err != nil {
			log.Err(err).Msg("Could not store resume context initial prompt from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt in db"})
			return
		}
		err = firebaseService.StoreToPromptHistory(userId, geminiResponse, "model")
		if err != nil {
			log.Err(err).Msg("Could not store resume context prompt response from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt response in db"})
			return
		}

		// Return response for request
		c.JSON(http.StatusOK, gin.H{"text": extractedResumeText})
	}
}

func UpdateResume(c *gin.Context) {
	// Bind request body to UpdateResumeRequest struct
	var updateResumeRequest models.UpdateResumeRequest
	if err := c.ShouldBindJSON(&updateResumeRequest); err != nil {
		// Bad Request for non-JSON body
		log.Error().Msgf("Received UpdateResume request with non-JSON body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if required fields are present
	if updateResumeRequest.Experience == "" || updateResumeRequest.UserID == "" {
		log.Error().Msgf("Received UpdateResume request with malformed JSON body: %v", updateResumeRequest)
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Missing or empty required field"})
		return
	}

	log.Info().Msgf("Received UpdateResume request with body: %v", updateResumeRequest)

	c.JSON(http.StatusOK, &models.ErrorResponse{Message: "Experience saved successfully"})
}

func GetResume(firebaseService *services.FirebaseService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Source userId from path param
		userId := c.Param("userId")
		log.Info().Msgf("Processing get resume request for user: %s", userId)

		resume, err := firebaseService.GetResume(userId)
		if err != nil {
			log.Err(err).Msgf("Could not fetch resume for user: %s", userId)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Something failed while getting data"})
			return
		}

		promptHistory, err := firebaseService.GetResumePromptHistory(userId)
		if err != nil {
			log.Err(err).Msgf("Could not fetch resume prompt history for user: %s", userId)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Something failed while getting data"})
			return
		}

		c.JSON(http.StatusOK, models.GetResumeResponse{
			Resume:        *resume,
			PromptHistory: promptHistory,
		})
	}
}
