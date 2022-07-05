package controllers

import (
	"systemMoniter-Server/auth"
	"systemMoniter-Server/common"
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
		logic.ResponseError(context, logic.ResponseCode(common.ErrMarshalJson.Code), common.ErrMarshalJson.Errord)
		context.Abort()
		return
	}
	user := models.User{
		User:     loginUser.User,
		Password: loginUser.Password,
	}
	if err := user.HashPassword(user.Password); err != nil {
		msg := "Create User " + user.User + " and Hash password error"
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrEncrypt.Code), common.ErrEncrypt.Errord)
		context.Abort()
		return
	}

	err := mysql.InsertUser(&user)
	if err != nil {
		msg := "Create User " + user.User + " error: " + err.Error()
		logic.ResponseError(context, logic.ResponseCode(common.ErrDuplicateUserFound.Code), common.ErrDuplicateUserFound.Errord)
		zap.L().Error(msg)
		context.Abort()
		return
	}
	msg := "Create User " + user.User + " , Uid:" + user.Id + " successfull "
	zap.L().Info(msg)
	logic.ResponseSuccess(context, gin.H{"userId": user.Id, "name": user.User})

}

func JsonLogin(context *gin.Context) {
	var loginUser models.LoginUser
	if err := context.ShouldBindJSON(&loginUser); err != nil {
		msg := "Login User " + loginUser.User + "error: "
		zap.L().Error(msg, zap.Error(err))
		logic.ResponseError(context, logic.ResponseCode(common.ErrMarshalJson.Code), common.ErrMarshalJson.Errord)
		context.Abort()
		return
	}
	var user models.User
	err := mysql.FindUser(loginUser.User, &user)
	if err != nil {
		msg := "Login User " + loginUser.User + " error: " + err.Error()
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrUserNotFound.Code), common.ErrUserNotFound.Errord)
		context.Abort()
		return
	}
	err = user.CheckPassword(loginUser.Password)
	//fmt.Println(err.Error())
	if err != nil {
		logic.ResponseError(context, logic.ResponseCode(common.ErrUsernameOrPasswordIncorrect.Code), common.ErrUsernameOrPasswordIncorrect.Errord)
		msg := "Check User " + user.User + "password error: " + err.Error()
		zap.L().Error(msg)
		//context.Abort()
		return
	}
	token, err := auth.GenerateJWT(loginUser.User)
	if err != nil {
		msg := "Login User " + loginUser.User + "and Create jwt error" + err.Error()
		zap.L().Error(msg)
		logic.ResponseError(context, logic.ResponseCode(common.ErrTokenSign.Code), common.ErrTokenSign.Errord)
		//context.Abort()
		return
	}
	msg := "Login User " + user.User + " , Uid:" + user.Id + " successfull "
	zap.L().Info(msg)
	logic.ResponseSuccess(context, gin.H{"userId": user.Id, "name": user.User, "token": token})

}
