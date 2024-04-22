package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func SetupMiddlewares(router *gin.Engine) {
	// RequestId middleware
	router.Use(requestid.New())

	// Logging middleware
	router.Use(
		logger.SetLogger(
			logger.WithLogger(func(ctx *gin.Context, logger zerolog.Logger) zerolog.Logger {
				return logger.With().Str("request_id", requestid.Get(ctx)).Logger()
			}),
		),
	)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin, X-Requested-With, Content-Type, Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
