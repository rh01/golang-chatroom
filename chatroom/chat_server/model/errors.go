package model

import "errors"

// 定义错误的类型
var (
	ErrUserNotExist  = errors.New("user not exist")
	ErrInvalidPasswd = errors.New("Passwd or username not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserExist     = errors.New("user exist")
)
