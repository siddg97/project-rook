package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/siddg97/project-rook/pkg/models"
	"net/http"
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
		c.JSON(http.StatusOK, models.PingResponse{Ping: "pong"})
	}
}
