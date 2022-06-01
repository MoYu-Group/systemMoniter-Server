package controllers

import (
	"net/http"
	"systemMoniter-Server/dao/mysql"
	"systemMoniter-Server/logic"
	"systemMoniter-Server/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// func FormRegisterUser(context *gin.Context) {
// 	var loginUser models.LoginUser
// 	if err := context.ShouldBind(&loginUser); err != nil {
// 		msg := "Create User " + loginUser.User + "error: "
// 		zap.L().Error(msg, zap.Error(err))
// 		logic.ResponseError(context, http.StatusBadRequest, err)
// 		context.Abort()
// 		return
// 	}
// 	user := models.User{
// 		User:     loginUser.User,
// 		Password: loginUser.Password,
// 	}
// 	if err := user.HashPassword(user.Password); err != nil {
// 		msg := "Create User " + user.User + "error: "
// 		zap.L().Error(msg, zap.Error(err))
// 		logic.ResponseError(context, http.StatusInternalServerError, err)
// 		context.Abort()
// 		return
// 	}

// 	err := mysql.InsertUser(&user)
// 	if err != nil {
// 		logic.ResponseError(context, http.StatusInternalServerError, err)
// 		msg := "Create User " + user.User + "error: "
// 		zap.L().Fatal(msg, zap.Error(err))
// 		context.Abort()
// 		return
// 	}
// 	msg := "Create User " + user.User + " , Uid:" + user.Id + " successfull "
// 	zap.L().Info(msg)
// 	logic.ResponseSuccess(context, gin.H{"userId": user.ID, "name": user.User})

// }

func JsonRegisterUser(context *gin.Context) {
	var loginUser models.LoginUser
	if err := context.ShouldBindJSON(&loginUser); err != nil {
		msg := "Create User " + loginUser.User + "error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, http.StatusBadRequest, err)
		context.Abort()
		return
	}
	user := models.User{
		User:     loginUser.User,
		Password: loginUser.Password,
	}
	if err := user.HashPassword(user.Password); err != nil {
		msg := "Create User " + user.User + "error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, http.StatusInternalServerError, err)
		context.Abort()
		return
	}

	err := mysql.InsertUser(&user)
	if err != nil {
		logic.ResponseError(context, http.StatusInternalServerError, err)
		msg := "Create User " + user.User + "error: "
		zap.L().Fatal(msg, zap.Error(err))
		context.Abort()
		return
	}
	msg := "Create User " + user.User + " , Uid:" + user.Id + " successfull "
	zap.L().Info(msg)
	logic.ResponseSuccess(context, gin.H{"userId": user.Id, "name": user.User})

}
