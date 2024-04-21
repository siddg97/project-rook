package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/siddg97/project-rook/models"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.Pong{Message: "pong"})
}
