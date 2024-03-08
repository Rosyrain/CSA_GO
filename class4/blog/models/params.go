package models

// 定义请求的参数结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp  注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin  登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// VoteData  投票数据
type ParamVoteData struct {
	//userID 从请求中获取当前的用户
	PostID    string `json:"post_id" binding:"required"`              //帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票1，反对票-1，取消图片（0）
}

// ParamPostList  获取帖子列表query string 参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id" example:"1"` //	可以为空
	Page        int64  `json:"page" form:"page" example:"1"`                 // 页码
	Size        int64  `json:"size" form:"size" example:"10"`                // 每页数量
	Order       string `json:"order" form:"order" example:"score"`           // 排序依据
}

// ParamCommunityPostList  获取社区列表query string 参数
//type ParamCommunityPostList struct {
//*ParamPostList
//CommunityID int64  `json:"community_id" form:"community_id"`   //	可以为空
//}
