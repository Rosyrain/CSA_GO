package controller

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// SignUpHandler 创建用户

// @Summary 创建用户
// @Description 创建新用户接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamSignUp true "用户注册参数"
// @Success 1000 {object} _ResponseSuccess
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	//1.获取参数并进行校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))

		return
	}

	fmt.Println(p)

	//业务处理
	if err := logic.SignUp(p); err != nil {
		//
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回相应
	ResponseSuccess(c, nil)

}

// LoginHandler 实现登录的处理函数
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamLogin true "用户登录参数"
// @Success 1000 {object} _ResponseSuccess
func LoginHandler(c *gin.Context) {
	//1.获取参数并进行校验
	p := new(models.ParamLogin)

	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	//2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}

	//3.返回参数
	ResponseSuccess(c, gin.H{
		"user_id":   user.UserID,
		"user_name": user.Username,
		"token":     user.Token,
	})
}
