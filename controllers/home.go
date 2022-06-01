package controllers

import (
	"systemMoniter-Server/logic"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	logic.ResponseSuccess(c, string("Welcome to use!"))
}
