package routes

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/siddg97/project-rook/handlers"
)

func SetupRoutes(ctx context.Context, router *gin.Engine, geminiClient *genai.Client) {
	versionOne := router.Group("/v1")
	{
		versionOne.GET("/story", handlers.GenerateStory(ctx, geminiClient))
	}

	router.GET("/healthcheck", handlers.HealthCheck)
	router.GET("/ping", handlers.Ping)
}
