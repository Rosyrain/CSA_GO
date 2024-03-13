package logic

import (
	"blog/dao/redis"
	"blog/models"
	"go.uber.org/zap"
	"strconv"
)

func PostForVote(userID int64, p *models.ParamVoteData) (err error) {
	zap.L().Debug("PostForVote",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))

}
