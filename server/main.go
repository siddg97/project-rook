package main

import (
	"context"
	"fmt"

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
	// Setup server config
	cfg, configErr := config.InitConfig()
	if configErr != nil {
		log.Fatal().Msgf("Failed to initialize server config due to fatal error %v", configErr)
		panic(configErr)
	}

	ctx := context.Background()

	// Setup Firebase
	firebaseConfigPath := "firebase_credentials.json"
	firestoreClient, firestoreErr := services.InitializeFirestore(ctx, firebaseConfigPath)
	if firestoreErr != nil {
		log.Fatal().Msgf("Failed to initialize Firestore due to fatal error %v", firestoreErr)
		panic(firestoreErr)
	}
	defer firestoreClient.Close()

	// Setup Gemini
	geminiClient, geminiErr := services.InitializeGemini(ctx, cfg.GeminiApiKey)
	if geminiErr != nil {
		panic(geminiErr)
	}
	defer geminiClient.Close()

	// Setup Cloud Vision
	_, visionErr := services.InitializeVision(ctx)
	if visionErr != nil {
		panic(visionErr)
	}

	// Initialize Gin
	gin.SetMode(cfg.LogLevel)
	router := gin.New()

	server := Server{
		geminiClient:    geminiClient,
		router:          router,
		firestoreClient: firestoreClient,
	}

	// Setup middlewares
	middlewares.SetupMiddlewares(router)

	// Setup routes
	routes.SetupRoutes(router)

	// Start server
	if err := server.router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Fatal().Msgf("Failed to run server: %v", err)
	}
}
