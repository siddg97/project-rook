package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type UpdateResumeRequest struct {
	UserID     string `json:"userId"`
	Experience string `json:"experience"`
}

func CreateResume(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Needs implementation"})
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
