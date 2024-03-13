package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPostListHandler 获取帖子列表的处理函数
// @Summary 获取帖子列表
// @Description 获取帖子列表接口，根据参数查询帖子列表
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object body models.ParamPostList true "查询参数"
// @Security ApiKeyAuth
// @Success 1000 {object} _ResponseSuccess

// @Router /posts [post]
func GetPostListHandler(c *gin.Context) {
	//初始胡结构体参数
	p := models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	//1.获取参数
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("GetPostListHandler with invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//2.业务处理
	data, err := logic.GetPostList(&p)
	if err != nil {
		zap.L().Error("logic.GetPostList(&p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, data)
}
