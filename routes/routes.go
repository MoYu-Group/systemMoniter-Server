package routes

import (
	"systemMoniter-Server/controllers"
	"systemMoniter-Server/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", controllers.Home)
	api := r.Group("/api")
	{
		api.POST("/user/register", controllers.JsonRegisterUser)
	}
	return r
}
