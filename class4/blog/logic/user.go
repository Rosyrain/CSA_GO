package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	snowflake "blog/pkg/snowflask"
)

// SignUp 根据controller层的参数进行  注册  业务处理并返回数据或错误
func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//生成UID
	userID := snowflake.GenID()

	//构建一个User实例对象
	u := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	//保存进数据库
	return mysql.InsertUser(&u)

}

// Login 根据controller层的参数进行  登录  业务处理并返回数据或错误
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	if err := mysql.Login(user); err != nil {
		return nil, err
	}

	//生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
