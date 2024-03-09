package logic

import (
	"blog/dao/mysql"
	"blog/dao/redis"
	"blog/models"
	snowflake "blog/pkg/snowflask"
)

// CreatePost 完成帖子创建的业务处理
func CreatePost(p *models.Post) (err error) {
	//1.生成post_id
	p.ID = snowflake.GenID()

	//2.保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID)
	if err != nil {
		return err
	}

	//3.返回
	return

}

// DeletePost 完成帖子删除的业务处理
func DeletePost(pid, uid int64) (err error) {
	//1.判断当前帖子是否存在
	if err = mysql.CheckPostExist(pid); err != nil {
		return err
	}

	//2.判断当前帖子的创建者是否为当前用户
	if err = mysql.CheckPostCreator(uid, pid); err != nil {
		return err
	}

	//3.删除帖子
	if err = mysql.DeletePost(pid); err != nil {
		return err
	}

	//4.返回
	return
}
