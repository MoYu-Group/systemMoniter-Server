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
	api.POST("/user/login", controllers.JsonLogin)
	api.GET("/node/allStatus", controllers.GetAllNodeStatus)
	api.Use(logic.Auth())
	{
		//api.GET("/node/local", controllers.Local)
		api.POST("/node/saveStatus", controllers.JsonStatus)
		api.POST("/node/register", controllers.JsonRegisterNode)
		api.POST("/user/register", controllers.JsonRegisterUser)
		api.GET("/auth", controllers.AuthTestPassed)
	}
	return r
}
