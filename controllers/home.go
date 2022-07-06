package controllers

import (
	"systemMoniter-Server/common/local"
	"systemMoniter-Server/logic"
	"systemMoniter-Server/models"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	logic.ResponseSuccess(c, string("Welcome to use!"))
}

func Local(c *gin.Context) {
	info := models.NewDefaultStatus()
	info.Load1 = local.GetBasic.Load1
	info.Load5 = local.GetBasic.Load5
	info.Load15 = local.GetBasic.Load15
	info.Thread = local.GetBasic.Thread
	info.Process = local.GetBasic.Process
	info.NetworkTx = local.GetNetSpeed.Avgtx
	info.NetworkRx = local.GetNetSpeed.Avgrx
	info.NetworkIn = uint64(local.GetNetSpeed.Nettx)
	info.NetworkOut = uint64(local.GetNetSpeed.Netrx)
	info.Ping10010 = local.PingValue.Ping10010
	info.Ping10086 = local.PingValue.Ping10086
	info.Ping189 = local.PingValue.Ping189
	info.Time10010 = local.PingValue.Time10010
	info.Time10086 = local.PingValue.Time10086
	info.Time189 = local.PingValue.Time189
	info.TCP = local.GetBasic.TCP
	info.UDP = local.GetBasic.UDP
	info.CPU = local.GetBasic.CPU
	info.MemoryTotal = local.GetBasic.MemoryTotal
	info.MemoryUsed = local.GetBasic.MemoryUsed
	info.SwapTotal = local.GetBasic.SwapTotal
	info.SwapUsed = local.GetBasic.SwapUsed
	info.Uptime = local.GetBasic.Uptime
	info.HddTotal = local.GetBasic.HddTotal
	info.HddUsed = local.GetBasic.HddUsed
	info.IpStatus = local.PingValue.IpStatus
	logic.ResponseSuccess(c, info)

}

func AuthTestPassed(c *gin.Context) {
	logic.ResponseSuccess(c, string("Auth test passed!"))
}
