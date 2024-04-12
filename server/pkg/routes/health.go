package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/siddg97/project-rook/pkg/models"
	"google.golang.org/api/option"
)

func ConfigureChecks(router *gin.Engine) {
	checkRoutes := router.Group("/checks")
	{
		checkRoutes.GET("/health", healthCheck())
		checkRoutes.GET("/ping", ping())
	}
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNoContent, nil)
	}
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBwVRHg5_GcRKqA-m252slgcTHtyqvQRsU"))
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}
		defer client.Close()
		model := client.GenerativeModel("gemini-pro")

		prompt_text := genai.Text("Write a story about a magic backpack.")
		resp, err := model.GenerateContent(ctx, prompt_text)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response := formatResponse(resp)
		log.Println(response)

		c.JSON(http.StatusOK, models.PingResponse{Ping: response})
	}
}

func formatResponse(resp *genai.GenerateContentResponse) string {
	var formattedContent strings.Builder
	if resp != nil && resp.Candidates != nil {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formattedContent.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formattedContent.String()
}
