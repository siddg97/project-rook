package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/siddg97/project-rook/models"
	"github.com/siddg97/project-rook/services"

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

		createResumeRequest := models.CreateResumeRequest{
			Resume: requestFile,
		}

		if createResumeRequest.Resume == nil {
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Resume file is not attached to the CreateResume request"})
			return
		}

		openedResume, _ := createResumeRequest.Resume.Open()
		filename := requestFile.Filename
		log.Info().Msgf("Opened file %v in CreateResumeRequest", filename)

		ResumePdfContent, err := io.ReadAll(openedResume)
		if err != nil {
			log.Error().Msgf("Error to read file %s: %v", filename, err)
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Failed to read attached resume file. Ensure that the file type is PDF"})
			return
		}

		// Extract text from PDF file via GCP Vision APIs
		extractedResumeText, err := visionService.ExtractTextFromPdf(ResumePdfContent)
		if err != nil {
			log.Error().Msgf("Encountered error when trying to extract text from PDF file: %v", err)
			c.JSON(http.StatusInternalServerError, &models.ErrorResponse{Message: "Failed to extract text from resume"})
			return
		}
		log.Info().Msgf("Successfully extracted %v characters of text for attached resume for userId: %v", len(extractedResumeText), userId)

		// Store resume text for user
		err = firebaseService.StoreNewResume(userId, extractedResumeText)
		if err != nil {
			log.Err(err).Msg("Could not store resume to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume in db"})
			return
		}

		// Setup context with initial prompt to gemini
		initialContextPrompt := models.GetInitialResumeCreationPrompt(userId, extractedResumeText)
		summarizedResumeDetails, err := geminiService.PromptGemini(initialContextPrompt)
		if err != nil {
			log.Err(err).Msg("Failed to save context for resume via gemini prompt")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error prompting Gemini for initial resume context"})
			return
		}
		log.Info().Msgf("%s", summarizedResumeDetails)

		// Ensure it is the correct JSON format
		var resumeDetails models.ModelResumeDetails
		err = json.Unmarshal([]byte(summarizedResumeDetails), &resumeDetails)
		if err != nil {
			log.Err(err).Msgf("Gemini model response failed to conform to the expected ResumeDetails struct: %v", summarizedResumeDetails)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error found in Gemini generated response"})
			return
		}

		// Store gemini prompt and response to prompt history in db
		err = firebaseService.StoreToPromptHistory(userId, initialContextPrompt, "user")
		if err != nil {
			log.Err(err).Msg("Could not store resume context initial prompt from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt in db"})
			return
		}
		err = firebaseService.StoreToPromptHistory(userId, summarizedResumeDetails, "model")
		if err != nil {
			log.Err(err).Msg("Could not store resume context prompt response from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt response in db"})
			return
		}

		// Return response for request
		c.JSON(http.StatusOK, models.CreateResumeResponse{
			Resume:        extractedResumeText,
			ResumeDetails: resumeDetails,
		})
	}
}

func UpdateResume(firebaseService *services.FirebaseService, geminiService *services.GeminiService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Source userId from path param
		userId := c.Param("userId")
		log.Info().Msgf("Processing resume update request for user: %s", userId)

		// Bind request body to UpdateResumeRequest struct
		var updateResumeRequest models.UpdateResumeRequest
		if err := c.ShouldBindJSON(&updateResumeRequest); err != nil {
			// Bad Request for non-JSON body
			log.Error().Msgf("Received UpdateResume request with non-JSON body")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
			return
		}

		// Check if required fields are present
		if updateResumeRequest.Experience == "" {
			log.Error().Msgf("Received UpdateResume request with malformed JSON body: %v", updateResumeRequest)
			c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Missing or empty required field"})
			return
		}

		log.Info().Msgf("Received UpdateResume request with experience: %v", updateResumeRequest.Experience)

		// Retrieve prompt history for the user
		promptHistory, err := firebaseService.GetResumePromptHistory(userId)
		if err != nil {
			log.Err(err).Msgf("Could not fetch resume prompt history for user: %s", userId)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Something failed while getting data"})
			return
		}
		if len(promptHistory) < 2 {
			log.Error().Msgf("Expected at least 2 resume prompts before update request can be processed, but found %s prompt(s)", len(promptHistory))
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: fmt.Sprintf("Expected at least 2 resume prompts before update request can be processed, but found %s prompt(s)", len(promptHistory))})
			return
		}

		// Update the resume with the experience prompt
		addExperiencePrompt := models.AddExperiencePrompt(userId, updateResumeRequest.Experience)
		newResumeDetails, err := geminiService.PromptGeminiWithHistory(promptHistory, addExperiencePrompt)
		if err != nil {
			log.Err(err).Msg("Failed to save new experience for user via gemini prompt")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error prompting Gemini for updating the user's new experience"})
			return
		}

		log.Info().Msgf("%s", newResumeDetails)

		// Ensure it is the correct JSON format
		var resumeDetails models.ModelResumeDetails
		err = json.Unmarshal([]byte(newResumeDetails), &resumeDetails)
		if err != nil {
			log.Err(err).Msgf("Gemini model response failed to conform to the expected ResumeDetails struct: %v", newResumeDetails)
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error found in Gemini generated response"})
			return
		}

		log.Info().Msgf("Received response from Gemini after adding new experience: %v", resumeDetails)

		// Store gemini prompt and response to prompt history in db
		err = firebaseService.StoreToPromptHistory(userId, addExperiencePrompt, "user")
		if err != nil {
			log.Err(err).Msg("Could not store resume context initial prompt from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt in db"})
			return
		}
		err = firebaseService.StoreToPromptHistory(userId, newResumeDetails, "model")
		if err != nil {
			log.Err(err).Msg("Could not store resume context prompt response from gemini to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume context prompt response in db"})
			return
		}

		// Store new resume in db
		err = firebaseService.StoreNewResume(userId, newResumeDetails)
		if err != nil {
			log.Err(err).Msg("Could not store updated resume to firebase")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error storing resume in db"})
			return
		}

		c.JSON(http.StatusOK, models.UpdateResumeResponse{
			Resume:        newResumeDetails,
			ResumeDetails: resumeDetails,
		})
	}
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
