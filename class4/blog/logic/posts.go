package logic

import (
	"blog/dao/redis"
	"blog/models"
	"go.uber.org/zap"
)

//type Post struct {
//	ID         int64     `json:"id,string" db:"post_id" example:""`
//	AuthorID   int64     `json:"author_id,string" db:"author_id" example:""`
//	Title      string    `json:"title" db:"title" binding:"required" example:""`
//	Content    string    `json:"content" db:"content" binding:"required" example:""`
//	CreateTime time.Time `json:"create_time" db:"create_time"`
//}
//
//type ApiPostDetail struct {
//	AuthorName string `json:"author_name" example:"rosyrain"`
//	VoteNumber int64  `json:"vote_number" example:"10"`
//	*Post             //嵌入式帖子结构体
//}

func GetPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//1.根据查询形式,在redis中拿到对应的id列表
	ids, err := redis.GetPostListIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostListIDsInOrder success but return 0 data")
		return
	}
	zap.L().Debug("GetPostListByTime", zap.Any("ids", ids))

	//2.根据id去mysql数据库中查询帖子详细详细
	posts, err := mysql.GetPostListByIDs(ids)

	//将帖子的作者及分区去查询出来填充至帖子中
	data = make([]*models.ApiPostDetail, 0, len(ids))
	//提前将查询好每篇帖子的投票数据
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return nil, err
	}

	zap.L().Debug("GetPostListByTime ", zap.Any("vote data", voteData))

	//3.组合数据
	for index, post := range posts {
		//根据作者id查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("authorID", post.AuthorID),
				zap.Error(err))
			continue
		}
		//组合数据
		postDetail := &models.ApiPostDetail{
			AuthorName: user.Username,
			VoteNumber: voteData[index],
			Post:       post,
		}
		//追加到data中
		data = append(data, postDetail)
	}
	return
}
