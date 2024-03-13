package models

import "time"

//内存对齐概念

type Post struct {
	ID         int64     `json:"id,string" db:"post_id" example:""`
	AuthorID   int64     `json:"author_id,string" db:"author_id" example:""`
	Title      string    `json:"title" db:"title" binding:"required" example:""`
	Content    string    `json:"content" db:"content" binding:"required" example:""`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name" example:"rosyrain"`
	VoteNumber int64  `json:"vote_number" example:"10"`
	*Post             //嵌入式帖子结构体
}
