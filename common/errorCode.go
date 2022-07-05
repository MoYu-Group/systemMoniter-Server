package common

import (
	"errors"
)

// 定义错误码
type Errno struct {
	Code    int
	Message string
	Errord  error // 保存内部错误信息
}

func (err Errno) Error() string {
	return err.Message
}

var (
	OK = &Errno{Errord: errors.New(""), Code: 0, Message: "OK"}
	// 系统错误, 前缀为 100
	InternalServerError = &Errno{Errord: errors.New("InternalServerError"), Code: 10001, Message: "内部服务器错误"}
	ErrBind             = &Errno{Errord: errors.New("ErrBind"), Code: 10002, Message: "请求参数错误"}
	ErrTokenSign        = &Errno{Errord: errors.New("ErrTokenSign"), Code: 10003, Message: "签名 jwt 时发生错误"}
	ErrEncrypt          = &Errno{Errord: errors.New("ErrEncrypt"), Code: 10005, Message: "加密用户密码时发生错误"}
	ErrMarshalJson      = &Errno{Errord: errors.New("ErrMarshalJson"), Code: 10005, Message: "解析 Json 错误"}

	// 数据库错误, 前缀为 201
	ErrDatabase = &Errno{Errord: errors.New("ErrDatabase"), Code: 20100, Message: "数据库错误"}
	ErrFill     = &Errno{Errord: errors.New("ErrFill"), Code: 20101, Message: "从数据库填充 struct 时发生错误"}

	// 认证错误, 前缀是 202
	ErrValidation   = &Errno{Errord: errors.New("ErrValidation"), Code: 20201, Message: "验证失败"}
	ErrTokenInvalid = &Errno{Errord: errors.New("ErrTokenInvalid"), Code: 20202, Message: "jwt 是无效的"}
	ErrTokenExpired = &Errno{Errord: errors.New("ErrTokenExpired"), Code: 20202, Message: "jwt 是过期的"}
	// 用户错误, 前缀为 203
	ErrUserNotFound                = &Errno{Errord: errors.New("ErrUserNotFound"), Code: 20301, Message: "用户没找到"}
	ErrUsernameOrPasswordIncorrect = &Errno{Errord: errors.New("ErrUsernameOrPasswordIncorrect"), Code: 20302, Message: "用户名或密码错误"}
	ErrDuplicateUserFound          = &Errno{Errord: errors.New("ErrDuplicateUserFound"), Code: 20303, Message: "重复用户"}
	ErrCreateUser                  = &Errno{Errord: errors.New("ErrCreateUser"), Code: 20304, Message: "创建用户失败"}
)
