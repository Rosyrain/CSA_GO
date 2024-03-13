package redis

//redis key

//redis key尽量使用命名空间的方式，方便查询和拆分

const (
	KeyPrefix          = "csa_blog:"
	KeyPostTime        = "post:time:"  //zset;帖子及发帖时间
	KeyPostScore       = "post:score:" //zset;帖子及投票分数
	KeyPostVotedPrefix = "post:voted:" //zset;记录用户投票类型；参数是post_id

	//KeyCommunityPrefix = "community:" //set:保存每个分区下的id
)

// GetRedisKey  返回redis key加上前缀
func GetRedisKey(key string) string {
	return KeyPrefix + key
}
