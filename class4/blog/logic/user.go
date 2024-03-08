package logic

import (
	"blog/models"
	"blog/pkg/jwt"
)

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户是否存在

}

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
}
