package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 实现帖子的创建
func CreatePostHandler(c *gin.Context) {
	//1.参数验证

	p := new(models.Post)

	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug(" c.ShouldBindJSON(p) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}

	//查看当前用户的id
	userId, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	p.AuthorID = userId

	//2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)

}

func DeletePostHandler(c *gin.Context) {
	//1.参数校验
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("delete post with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//获取当前用户uid判断当前用户是否为帖子的创建用户
	uid, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	//2.业务处理
	if err := logic.DeletePost(pid, uid); err != nil {
		zap.L().Error("logic.DeletePost(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回响应
	ResponseSuccess(c, nil)
}
