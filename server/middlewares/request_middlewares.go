package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/constants"
	"github.com/siddg97/project-rook/models"
	"net/http"
	"strings"
)

func ValidateCreateResumeRequest(c *gin.Context) {
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
	if requestFile == nil {
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "no resume file attached"})
		return
	}

	fileHasPdfExtension := strings.HasSuffix(requestFile.Filename, ".pdf")
	if !fileHasPdfExtension {
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Uploaded file is not a .pdf file"})
		return
	}

	createResumeRequest := &models.CreateResumeRequest{
		Resume: requestFile,
		UserId: userId,
	}
	c.Set(constants.CreateResumeRequestContextKey, createResumeRequest)

	log.Info().Msgf("Validated create resume request for user %s", userId)
	c.Next()
}

func ValidateUpdateResumeRequest(c *gin.Context) {
	// Source userId from path param
	userId := c.Param("userId")
	log.Info().Msgf("Processing resume update request for user: %s", userId)

	// Bind request body to UpdateResumeRequest struct
	updateResumeRequest := &models.UpdateResumeRequest{}
	if err := c.ShouldBindJSON(updateResumeRequest); err != nil {
		// Bad Request for non-JSON body
		log.Error().Msgf("Received invalid request with non-JSON body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Check if required fields are present
	if updateResumeRequest.Experience == "" {
		log.Error().Msgf("update resume request with no experience: %v", updateResumeRequest)
		c.JSON(http.StatusBadRequest, &models.ErrorResponse{Message: "Missing or empty required field"})
		return
	}

	updateResumeRequest.UserId = userId
	c.Set(constants.UpdateResumeRequestContextKey, updateResumeRequest)

	log.Info().Msgf("Validated update resume request for user %s with experience: %v", updateResumeRequest.UserId, updateResumeRequest.Experience)
	c.Next()
}
