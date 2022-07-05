package logic

// 公共返回封装

import (
	"encoding/json"
	"systemMoniter-Server/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	SuccessCode ResponseCode = 200
)

type ResponseCode int

type Response struct {
	ErrorCode ResponseCode `json:"error"`
	ErrorMsg  string       `json:"error_msg"`
	Data      interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	resp := &Response{ErrorCode: code, ErrorMsg: err.Error(), Data: ""}
	c.JSON(200, resp)
	response, err := json.Marshal(resp)
	if err != nil {
		zap.L().Error("Json Marshal error: ", zap.Error(err))
		return
	}
	c.Set("response", string(response))
	//c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{ErrorCode: ResponseCode(common.OK.Code), ErrorMsg: "", Data: data}
	c.JSON(200, resp)
	response, err := json.Marshal(resp)
	if err != nil {
		zap.L().Error("Json Marshal error: ", zap.Error(err))
		return
	}
	c.Set("response", string(response))
}
