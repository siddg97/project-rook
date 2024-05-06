package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siddg97/project-rook/handlers"
	"github.com/siddg97/project-rook/middlewares"
	"github.com/siddg97/project-rook/services"
)

func SetupRoutes(router *gin.Engine, visionService *services.VisionService, firebaseService *services.FirebaseService, geminiService *services.GeminiService) {
	router.GET("/ping", handlers.Ping)

	versionOne := router.Group("/v1")
	{
		versionOne.PUT(
			"/:userId/resume",
			middlewares.ValidateCreateResumeRequest,
			handlers.CreateResume(visionService, firebaseService, geminiService),
		)
		versionOne.POST(
			"/:userId/resume",
			middlewares.ValidateUpdateResumeRequest,
			handlers.UpdateResume(firebaseService, geminiService),
		)
		versionOne.GET("/:userId/resume", handlers.GetResume(firebaseService))
		versionOne.GET("/:userId", handlers.GetUserSummary(firebaseService, geminiService))
	}

}
