package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassowrd = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
	ErrorPostExist       = errors.New("帖子已存在")
	ErrorPostNotExist    = errors.New("帖子不存在")
	ErrorPostNotCreator  = errors.New("非帖子创建用户")
)
