package models

type User struct {
	UserID   int64  `db:"user_id" json:"user_id,string" `
	Username string `db:"username" example:"rosyrain"`
	Password string `db:"password" example:"abc123456"`
	Token    string `json:"token"`
}
