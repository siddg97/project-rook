package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siddg97/project-rook/handlers"
	"github.com/siddg97/project-rook/services"
)

func SetupRoutes(router *gin.Engine, visionService *services.VisionService, firebaseService *services.FirebaseService, geminiService *services.GeminiService) {
	router.GET("/ping", handlers.Ping)

	versionOne := router.Group("/v1")
	{
		versionOne.GET("/story", handlers.GenerateStory(geminiService))

		/**
		1. Receives file from frontend
		2. Extracts text (see if text extraction can be done in frontend)
		3. Saves resume to prompt history collection
		4. Returns 2XX
		*/
		versionOne.PUT("/:userId/resume", handlers.CreateResume(visionService, firebaseService))

		/**
		1. Receives a skill/experience update intent
		2. Loads up prompt history from before for user
		3. sets up chat history for gemini
		4. Process incoming update intent into prompt to send gemini
		5. Sends gemini prompt along with existing history
		6. Parse prompt response
		7. Updates the prompt history collection of the user
		8. Send log of updated resume section back to client
		*/
		versionOne.POST("/resume", handlers.UpdateResume)

		/**
		1. Load up prompt history so far from db
		2. Setup chat session with gemini
		3. Send prompt to summarize resume being tracked so far (along with chat history)
		4. Parse the response from gemini
		5. Upload to bucket (optionally)
		6. Return summary to client
		*/
		versionOne.GET("/:userId/resume", handlers.GetResume(firebaseService))

		/**
		Returns payload that can be used to determine if user has a resume or not
		*/
		//versionOne.GET("/user")
	}

}
