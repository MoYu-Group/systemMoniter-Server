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
	r.Use(logic.Cors())
	//r.GET("/", controllers.Home)
	r.Static("/css","./static/css")
	r.Static("/img","./static/img")
	r.Static("/js","./static/js")
	r.LoadHTMLGlob("./static/*.html")
	r.GET("/",func (c *gin.Context)  {
		c.HTML(200,"index.html",nil)
	})



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
