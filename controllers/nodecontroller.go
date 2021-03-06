package controllers

import (
	"systemMoniter-Server/common"
	"systemMoniter-Server/common/local"
	"systemMoniter-Server/dao/mysql"
	"systemMoniter-Server/logic"
	"systemMoniter-Server/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetAllNodeStatus(c *gin.Context) {
	err, servers := mysql.FindAllNodeStatus()
	if err != nil {
		msg := "Find all node status error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(c, logic.ResponseCode(common.ErrFindAllNodeInfo.Code), common.ErrFindAllNodeInfo.Errord)
		c.Abort()
		return
	}

	// for _, i := range servers {
	// 	cpu := i.CPU
	// 	i.CPU = float64(math.Ceil(cpu))
	// 	// uptime := int(i.Uptime)
	// 	// i.Uptime = fmt.Sprintln("%d天%d小时%d分%d秒\n",
	// 	// 	uptime/60/60/24%365,
	// 	// 	uptime/60/60%24,
	// 	// 	uptime/60%60,
	// 	// 	uptime%60)

	// }

	c.JSON(200, gin.H{"servers": servers, "updated": time.Now().Unix()})
}

func Local(c *gin.Context) {
	info := models.NewDefaultStatus()
	info.Load1 = local.GetBasic.Load1
	info.Load5 = local.GetBasic.Load5
	info.Load15 = local.GetBasic.Load15
	info.ThreadCount = local.GetBasic.Thread
	info.ProcessCount = local.GetBasic.Process
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
	info.TCPCount = local.GetBasic.TCP
	info.UDPCount = local.GetBasic.UDP
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

func JsonRegisterNode(context *gin.Context) {
	var nodeData models.NodeData
	if err := context.ShouldBindJSON(&nodeData); err != nil {
		msg := "Create Node " + nodeData.Name + "error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrCreateNode.Code), common.ErrCreateNode.Errord)
		context.Abort()
		return
	}
	var user models.User
	err := mysql.FindUserByUid(nodeData.Uid, &user)
	if err != nil && err.Error() == "No user find" {
		msg := "Create Node " + nodeData.Name + "error: no user find"
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrUserNotFound.Code), common.ErrUserNotFound.Errord)
		context.Abort()
		return
	}
	node := models.Node{
		Name:     nodeData.Name,
		Uid:      nodeData.Uid,
		Type:     nodeData.Type,
		Host:     nodeData.Host,
		Location: nodeData.Location,
		Custom:   nodeData.Custom,
		Disabled: false,
	}
	err2 := mysql.InsertNode(&node)
	if err2 != nil && err2.Error() == "Duplicate node find" {
		msg := "Create Node " + node.Name + " error: Duplicate node find"
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrDuplicateNodeFound.Code), common.ErrDuplicateNodeFound.Errord)
		context.Abort()
		return
	} else if err2 != nil {
		msg := "Create Node " + node.Name + " error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrCreateStatus.Code), common.ErrCreateStatus.Errord)
		context.Abort()
		return
	}
	msg := "Create Node " + node.Name + " , node id:" + node.Id + " successfull "
	zap.L().Info(msg)
	logic.ResponseSuccess(context, gin.H{"id": node.Id, "name": node.Name, "host": node.Host})
}
