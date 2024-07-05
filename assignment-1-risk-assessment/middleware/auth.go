package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	user = "admin"
	pass = "admin"
)

func AuthMiddleware(ctx *gin.Context) {
	username, password, ok := ctx.Request.BasicAuth()
	if username != user && password != pass && !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
