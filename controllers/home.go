package controllers

import (
	"net"
	"systemMoniter-Server/logic"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	ip := net.ParseIP(c.ClientIP())
	msg := string("Welcome to use! Your ip is ") + ip.String()
	logic.ResponseSuccess(c, msg)
}

func AuthTestPassed(c *gin.Context) {
	logic.ResponseSuccess(c, string("Auth test passed!"))
}
