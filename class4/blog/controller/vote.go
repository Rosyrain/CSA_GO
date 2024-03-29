package controller

import (
	"blog/logic"
	"blog/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//// VoteData  投票数据
//type VoteData struct {
//	//userID 从请求中获取当前的用户
//	PostID    int64 `json:"post_id,string"`   //帖子ID
//	Direction int   `json:"direction,string"` //赞成票1，反对票-1
//}

// PostVoteHandler 实现用户投票处理函数
// @Summary 用户投票
// @Description 用户投票接口
// @Tags 用户相关接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamVoteData true "用户投票参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /vote [post]
func PostVoteHandler(c *gin.Context) {
	//1.参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}

	//2.校验当前请求的用户ID
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	//3.业务处理
	if err := logic.PostForVote(userID, p); err != nil {
		zap.L().Error("logic.PostForVote failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//4.返回响应
	ResponseSuccess(c, nil)

}
