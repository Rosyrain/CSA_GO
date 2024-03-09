package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekINSeconds = 7 * 24 * 3600 //一周的秒数
	scorePerVote     = 432           //每一票是多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

// CreatePost 创建帖子时，创建对应记录
func CreatePost(postId int64) error {
	pipeline := client.TxPipeline()
	//帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTime), redis.Z{
		Score:  float64((time.Now().Unix())),
		Member: postId,
	})
	//帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	_, err := pipeline.Exec()
	return err
}

// VoteForPost 处理用户投票的数据更改
func VoteForPost(userID, postID string, value float64) error {
	//1.判断投票限制
	//去redis取出帖子发布时间
	postTime := client.ZScore(GetRedisKey(KeyPostTime), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekINSeconds {
		return ErrVoteTimeExpire
	}

	//2.更新分数
	//a.先查看当前用户的投票记录
	ov := client.ZScore(GetRedisKey(KeyPostVotedPrefix+postID), userID).Val()

	//b.更新：如果这一次投票的值和之前保持一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}

	//c.判断加减分
	var op float64 //用于标记符号，加分还是减分
	if value > ov {
		op = 1
	} else {
		op = -1
	}

	//d.计算差值
	diff := math.Abs(ov - value)

	pipeline := client.TxPipeline()
	//e.修改帖子得分
	pipeline.ZIncrBy(GetRedisKey(KeyPostScore), op*diff*scorePerVote, postID)

	//f.对帖子用户投票记录进行修改,如果value为0那么为取消投票，对应帖子删除uid记录;反之修改对应记录
	if value == 0 {
		pipeline.ZRem(GetRedisKey(KeyPostVotedPrefix+postID), userID)
	} else {
		pipeline.ZAdd(GetRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  value, //当前用户投的票型
			Member: userID,
		})
	}

	_, err := pipeline.Exec()
	return err
}
