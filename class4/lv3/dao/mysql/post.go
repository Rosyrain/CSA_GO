package mysql

import "blog/models"

// CreatePost 创建一条新的帖子信息到数据库当中
func CreatePost(p *models.Post) (err error) {
	sqlStr := `Insert into post(
        		post_id, title, content, author_id
                )values (?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID)
	return
}

func CheckPostExist(pid int64) (err error) {
	sqlStr := `select count(author_id) from post where post_id=?`
	var count int
	if err = db.Get(&count, sqlStr, pid); err != nil {
		return err
	}
	if count == 0 {
		return ErrorPostNotExist
	}
	return
}

func CheckPostCreator(uid, pid int64) (err error) {
	sqlStr := `select author_id from post where post_id=?`
	var author_id int64
	if err = db.Get(&author_id, sqlStr, pid); err != nil {
		return err
	}
	if author_id != uid {
		return ErrorPostNotCreator
	}
	return
}

// DeletePost 根据帖子id进行删除
func DeletePost(pid int64) (err error) {
	sqlStr := `DELETE FROM post WHERE post_id = ?`
	if _, err = db.Exec(sqlStr, pid); err != nil {
		return err
	}
	return
}
