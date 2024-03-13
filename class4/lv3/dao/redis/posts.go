package redis

import "blog/models"

func GetIDsFormKey(key string, page, size int64) ([]string, error) {
	//在Redis中，有序集合（Sorted Set）的查询下标是从0开始的，
	//并且是左闭右闭的范围。

	//1.确定始末
	start := (page - 1) * size
	end := start + size - 1

	//2.ZREVRANGE  按分数从大到小查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

// GetPostListIDsInOrder 按照时间顺序获取帖子IDs
func GetPostListIDsInOrder(p *models.ParamPostList) (ids []string, err error) {
	//从reids中获取ids

	//根据用户请求中携带的order参数确定要查询的redis key
	key := GetRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = GetRedisKey(KeyPostScore)
	}

	//2.确定查询的索引起始
	return GetIDsFormKey(key, p.Page, p.Size)
}
