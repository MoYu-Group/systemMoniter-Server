package controllers

import (
	"systemMoniter-Server/common"
	"systemMoniter-Server/dao/mysql"
	"systemMoniter-Server/logic"
	"systemMoniter-Server/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JsonStatus(context *gin.Context) {
	var status models.Status
	if err := context.ShouldBindJSON(&status); err != nil {
		msg := "Create Node " + status.Name + " status error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrCreateStatus.Code), common.ErrCreateStatus.Errord)
		context.Abort()
		return
	}
	if status.Name == "" || status.Host == "" {
		msg := "Create Node " + status.Name + " status error: lost node info"
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrLostNodeInfo.Code), common.ErrLostNodeInfo.Errord)
		context.Abort()
		return
	}
	isIPv4 := common.IsIPv4(context.ClientIP())
	isIPv6 := common.IsIPv6(context.ClientIP())
	info := models.Info{
		ProcessCount: status.ProcessCount,
		NetworkTx:    status.NetworkTx,
		NetworkRx:    status.NetworkRx,
		NetworkIn:    status.NetworkIn,
		NetworkOut:   status.NetworkOut,
		Ping10010:    status.Ping10010,
		Ping10086:    status.Ping10086,
		Ping189:      status.Ping189,
		Time10010:    status.Time10010,
		Time10086:    status.Time10086,
		Time189:      status.Time189,
		TCPCount:     status.TCPCount,
		UDPCount:     status.UDPCount,
		CPU:          status.CPU,
		MemoryTotal:  status.MemoryTotal,
		MemoryUsed:   status.MemoryUsed,
		SwapTotal:    status.SwapTotal,
		SwapUsed:     status.SwapUsed,
		Uptime:       status.Uptime,
		HddTotal:     status.HddTotal,
		HddUsed:      status.HddUsed,
		IpStatus:     status.IpStatus,
		Online4:      isIPv4,
		Online6:      isIPv6,
	}
	err := mysql.InsertInfoByNameAndHost(status.Name, status.Host, &info)
	if err != nil && err.Error() == "Node is disabled" {
		msg := "Create Node " + status.Name + " status error: node is disabled"
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrNodeIsDisabled.Code), common.ErrNodeIsDisabled.Errord)
		context.Abort()
		return
	} else if err != nil && err.Error() == "No node find" {
		msg := "Create Node " + status.Name + "status error: no node find"
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrNodeNotFound.Code), common.ErrNodeNotFound.Errord)
		context.Abort()
		return
	} else if err != nil {
		msg := "Create Node " + status.Name + "status error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrCreateStatus.Code), common.ErrCreateStatus.Errord)
		context.Abort()
		return
	}
	msg := "Create Node " + status.Name + " status successfull "
	zap.L().Info(msg)
	logic.ResponseSuccess(context, gin.H{"id": info.Id, "name": status.Name, "host": status.Host})
}
