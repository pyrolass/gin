package middlewares

import (
	"net/http"
	"test/common"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extract the token from the Authorization header
		tknStr := ctx.GetHeader("Authorization")

		println(tknStr)

		// Validate the token
		claims, err := common.VerifyToken(tknStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Token is valid, pass the claims to the next middleware
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
