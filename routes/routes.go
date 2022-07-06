package routes

import (
	"systemMoniter-Server/controllers"
	"systemMoniter-Server/logger"
	"systemMoniter-Server/logic"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", controllers.Home)
	api := r.Group("/api/")
	{
		api.POST("/local", controllers.Local)
		api.POST("/user/register", controllers.JsonRegisterUser)
		api.POST("/user/login", controllers.JsonLogin)
		api.Use(logic.Auth())
		{
			api.POST("/auth", controllers.AuthTestPassed)
		}

	}
	return r
}
