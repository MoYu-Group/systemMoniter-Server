package logic

import (
	"systemMoniter-Server/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{
				"error": -1,
				"msg":   "request does not contain an access token",
			})
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{
				"error": -1,
				"msg":   err.Error(),
			})
			context.Abort()
			return
		}

		zap.L().Info("Validate JWT Token Pass, Host:" + claims.Host)
		context.Next()
	}
}
