package handlers

import (
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
		initalContextPrompt := fmt.Sprintf(`
			This is the initial resume text extracted from a resume file uploaded by the user denoted by their id %s.
			The resume extracted text is bounded within ~~~.
			~~~
			%s
			~~~
			
			Please keep note of this initial resume state going forward and expect new work summaries to be provided from the user in the future. 
			Please always incorporate the new work summaries into the resume only if it is significant enough. 
			At this point, please summarize the work that the user has done in a JSON format like the following example:
			{
				"userId": "a-user-id",
				"skills": [
					{
						"name": "AWS Lambda",
						"yearsOfExperience": 4,
					},
					...
					...
				],
				"work-summaries": [
					{
						"title": "Software engineer",
						"startDate: "December 2020",
						"endDate": "Present",
						"company": "Amazon",
						"work": [
							"Developed an API for tracking return-to-office attendance",
							"Mentored incompetent engineers to be competent",
							...
							...
						]
					},
					{
						"title": "Software engineer",
						"startDate: "December 2019",
						"endDate": "December 2020",
						"company": "Microsoft",
						"work": [
							"Developed an API for tracking return-to-office attendance",
							"Mentored incompetent engineers to be competent",
							...
							...
						]
					}
					...
					...
				]
			}
		`, userId, extractedResumeText)
		geminiResponse, err := geminiService.PromptGemini(initalContextPrompt)
		if err != nil {
			log.Err(err).Msg("Failed to save context for resume via gemini prompt")
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "Error prompting gemini for initial resume context"})
			return
		}

		log.Info().Msgf("Received response from Gemini: %v", geminiResponse)

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
		c.Data(http.StatusOK, "application/json", []byte(geminiResponse))
	}
}

func UpdateResume(c *gin.Context) {
	// Source userId from path param
	userId := c.Param("userId")
	log.Info().Msgf("Processing resume creation request for user: %s", userId)

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
