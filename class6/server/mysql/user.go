package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"server/models"
)

const secret = "liweizhou.com"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func CreateUser(nickname, OPassword, mobile string) (err error) {
	//对密码进行加密
	password := encryptPassword(OPassword)
	// 执行sql语句入库
	sqlStr := `insert into user(nickname,mobile,password) values(?,?,?)`

	_, err = db.Exec(sqlStr, nickname, mobile, password)
	return err
}

func GetUserInfo(nickname string) (user *models.User, err error) {
	// 执行sql语句入库
	sqlStr := `select id,password,nickname,brithDay,gender from user where nickname= ?`
	user = new(models.User)
	err = db.Get(&user, sqlStr, nickname)
	return user, err
}

func GetUserByID(id int32) (user *models.User, err error) {
	// 执行sql语句入库
	sqlStr := `select id,password,nickname,brithDay,gender from user where id= ?`
	user = new(models.User)
	err = db.Get(&user, sqlStr, id)
	return user, err
}

func UpdateUserInfo(id int32, nickname string, birthday uint64, gender string) error {
	sqlStr := `UPDATE user SET nickname = ?, birthday = ?, gender = ? WHERE id = ?`
	_, err := db.Exec(sqlStr, nickname, birthday, gender, id)
	return err
}

func CheckPassWord(OPassword, password string) (bool, error) {
	return encryptPassword(OPassword) == password, nil
}
