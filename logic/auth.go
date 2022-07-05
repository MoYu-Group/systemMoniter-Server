package logic

import (
	"systemMoniter-Server/auth"
	"systemMoniter-Server/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			ResponseError(context, ResponseCode(common.ErrTokenInvalid.Code), common.ErrTokenInvalid.Errord)
			context.Abort()
			return
		}
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			if err.Error() == "couldn't parse claims" {
				ResponseError(context, ResponseCode(common.ErrValidation.Code), common.ErrValidation.Errord)
			} else if err.Error() == "token expired" {
				ResponseError(context, ResponseCode(common.ErrTokenExpired.Code), common.ErrTokenExpired.Errord)
			} else {
				ResponseError(context, ResponseCode(common.ErrTokenInvalid.Code), common.ErrTokenInvalid.Errord)
			}
			context.Abort()
			return
		}

		zap.L().Info("Validate JWT Token Pass, Username:" + claims.Username)
		context.Next()
	}
}
