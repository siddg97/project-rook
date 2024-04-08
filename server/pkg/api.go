package pkg

import (
	"context"
	"fmt"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/siddg97/project-rook/pkg/routes"
)

type ApiServer struct {
	logLevel string
	port     string
	Ctx      context.Context
	Router   *gin.Engine
}

func NewApiServer(ctx context.Context, ginLogLevel string, port string) *ApiServer {
	gin.SetMode(ginLogLevel)
	router := gin.New()

	return &ApiServer{
		Router:   router,
		Ctx:      ctx,
		port:     port,
		logLevel: ginLogLevel,
	}
}

func (s *ApiServer) Start() error {
	s.setupMiddlewares()
	s.setupRoutes()

	return s.Router.Run(fmt.Sprintf(":%s", s.port))
}

func (s *ApiServer) setupMiddlewares() {
	// RequestId middleware
	s.Router.Use(requestid.New())

	// Logging middleware
	s.Router.Use(
		logger.SetLogger(
			logger.WithLogger(func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
				return l.With().Str("request_id", requestid.Get(c)).Logger()
			}),
		),
	)
}

func (s *ApiServer) setupRoutes() {
	routes.ConfigureChecks(s.Router)
}
