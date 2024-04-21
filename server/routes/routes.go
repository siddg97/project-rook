package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siddg97/project-rook/handlers"
)

func SetupRoutes(router *gin.Engine) {
	versionOne := router.Group("/v1")
	{
		versionOne.GET("/story", handlers.GenerateStory)
		versionOne.PUT("/resume", handlers.CreateResume)
		versionOne.POST("/resume", handlers.UpdateResume)
		versionOne.GET("/resume", handlers.GetResume)
	}

	router.GET("/healthcheck", handlers.HealthCheck)
	router.GET("/ping", handlers.Ping)
}
