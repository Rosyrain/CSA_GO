package models

type User struct {
	ID       int32  `db:"id" json:"user_id,string"`
	Nickname string `db:"nickname"`
	Password string `db:"password"`
	Gender   string `json:"gender"`
	BirthDay uint64 `db:"birthday"`
}
