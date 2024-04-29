package main

import (
	"context"
	"fmt"
	"github.com/siddg97/project-rook/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/siddg97/project-rook/config"
	"github.com/siddg97/project-rook/middlewares"
	"github.com/siddg97/project-rook/routes"
)

type Server struct {
	ctx             context.Context
	router          *gin.Engine
	firebaseService *services.FirebaseService
	visionService   *services.VisionService
	geminiService   *services.GeminiService
}

func (s *Server) SetupServer() {
	middlewares.SetupMiddlewares(s.router)
	routes.SetupRoutes(s.router, s.visionService, s.firebaseService, s.geminiService)
}

func main() {
	ctx := context.Background()

	// Setup server config
	cfg, err := config.InitConfig(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to initialize server config due to fatal error")
		panic(err)
	}

	// Setup Firebase
	firebaseService, err := services.InitializeFirebase(ctx)
	if err != nil {
		log.Err(err).Msg("Failed to initialize firebase clients")
		panic(err)
	}
	defer firebaseService.FirestoreClient.Close()

	// Setup Gemini
	geminiService, err := services.InitializeGemini(ctx, cfg.GeminiApiKey)
	if err != nil {
		log.Err(err).Msg("Failed to initialize gemini client")
		panic(err)
	}
	defer geminiService.GeminiClient.Close()

	// Setup Cloud Vision
	visionService, visionErr := services.InitializeVision(ctx)
	if visionErr != nil {
		log.Err(err).Msg("Failed to initialize GCP Vision due to fatal error")
		panic(visionErr)
	}

	// Initialize Gin
	gin.SetMode(cfg.LogLevel)
	router := gin.Default()

	server := &Server{
		ctx:             ctx,
		router:          router,
		firebaseService: firebaseService,
		visionService:   visionService,
		geminiService:   geminiService,
	}
	server.SetupServer()

	// Start server
	if err := server.router.Run(fmt.Sprintf(":%s", cfg.Port)); err != nil {
		log.Err(err).Msgf("Failed to run server: %v", err)
	}
}
