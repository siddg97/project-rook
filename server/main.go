package main

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/config"
	"github.com/siddg97/project-rook/middlewares"
	"github.com/siddg97/project-rook/routes"
	"github.com/siddg97/project-rook/services"
)

type Server struct {
	router          *gin.Engine
	firestoreClient *firestore.Client
	geminiClient    *genai.Client
}

func main() {
	server := Server{}

	// Initialize Gin
	router := gin.New()
	server.router = router

	// Setup middlewares
	middlewares.SetupMiddlewares(router)

	// Setup server config
	cfg, configErr := config.InitConfig()
	if configErr != nil {
		log.Fatal().Msgf("Failed to initialize server config due to fatal error %v", configErr)
		panic(configErr)
	}
	gin.SetMode(cfg.LogLevel)

	ctx := context.Background()

	// Setup Firebase
	firebaseConfigPath := "firebase_credentials.json"
	firestoreClient, firestoreErr := services.InitializeFirestore(ctx, firebaseConfigPath)
	if firestoreErr != nil {
		log.Fatal().Msgf("Failed to initialize Firestore due to fatal error %v", firestoreErr)
		panic(firestoreErr)
	}
	server.firestoreClient = firestoreClient

	// Setup Gemini
	geminiKey := os.Getenv("GEMINI_KEY")
	geminiClient, geminiErr := services.InitializeGemini(ctx, geminiKey)
	if geminiErr != nil {
		panic(geminiErr)
	}
	server.geminiClient = geminiClient

	// Setup routes
	routes.SetupRoutes(ctx, router, geminiClient)

	// Start server
	if err := server.router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal().Msgf("Failed to run server: %v", err)
	}
}
