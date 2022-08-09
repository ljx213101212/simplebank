package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) hellowolrd(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}
