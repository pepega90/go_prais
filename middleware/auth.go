package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	var (
		user = "admin"
		pass = "admin"
	)

	username, password, ok := ctx.Request.BasicAuth()
	if username != user && password != pass && !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "kepo"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
