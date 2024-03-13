package mysql

import (
	"blog/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

// GetPostListByIDs  根据给点的id列表查询帖子数据
func GetPostListByIDs(ids []string) (data []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,create_time
			from post
			where post_id in (?)
			order by FIND_IN_SET(post_id,?)
			    `

	//https://www.liwenzhou.com/posts/Go/sqlx/
	//多条查询
	//1.组装请求,query为组装的sqlStr，args为对应sqlStr的参数
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	//2.发送请求拿到参数
	err = db.Select(&data, query, args...) //!!!!别忘记最后的 ...
	return
}
