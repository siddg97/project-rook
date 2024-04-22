package handlers

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/services"
	"google.golang.org/api/vision/v1"
)

type CreateResumeRequest struct {
	UserID string `json:"userId"`
}

type UpdateResumeRequest struct {
	UserID     string `json:"userId"`
	Experience string `json:"experience"`
}

type GetResumeRequest struct {
	UserID string `json:"userId"`
}

func CreateResume(c *gin.Context) {
	ctx := context.Background()

	requestFile, readFileErr := c.FormFile("file")
	if readFileErr != nil {
		log.Fatal().Msgf("Error open file from input")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get file from request"})
		return
	}

	openedFile, _ := requestFile.Open()

	filename := requestFile.Filename

	pdfContent, readPdfErr := io.ReadAll(openedFile)
	if readPdfErr != nil {
		log.Fatal().Msgf("Error to read file with filename: %v", filename)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read file"})
		return
	}

	visionClient := services.GetVisionService()
	textDetectionRequest := &vision.BatchAnnotateFilesRequest{
		Requests: []*vision.AnnotateFileRequest{
			{
				InputConfig: &vision.InputConfig{
					MimeType: "application/pdf",
					Content:  base64.StdEncoding.EncodeToString(pdfContent),
				},
				Features: []*vision.Feature{
					{Type: "DOCUMENT_TEXT_DETECTION"},
				},
			},
		},
	}

	textDetectionResponse, textDetectionCallErr := visionClient.Files.Annotate(textDetectionRequest).Context(ctx).Do()
	if textDetectionCallErr != nil {
		log.Fatal().Msgf("Error to call vision client to annotate files: %v", textDetectionCallErr.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to call Vision API to annotate file"})
		return
	}

	var extractedTextFromPdf string
	for _, page := range textDetectionResponse.Responses[0].Responses {
		extractedTextFromPdf += page.FullTextAnnotation.Text
	}

	c.JSON(http.StatusOK, gin.H{"text": extractedTextFromPdf})
}

func UpdateResume(c *gin.Context) {
	// Bind request body to UpdateResumeRequest struct
	var updateResumeRequest UpdateResumeRequest
	if err := c.BindJSON(&updateResumeRequest); err != nil {
		// Bad Request for non-JSON body
		log.Error().Msgf("Received UpdateResume request with non-JSON body")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if required fields are present
	if updateResumeRequest.Experience == "" || updateResumeRequest.UserID == "" {
		log.Error().Msgf("Received UpdateResume request with malformed JSON body: %v", updateResumeRequest)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or empty required field"})
		return
	}

	log.Info().Msgf("Received UpdateResume request with body: %v", updateResumeRequest)

	c.JSON(http.StatusOK, gin.H{"message": "Experience saved successfully"})
}

func GetResume(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Needs implementation"})
}
