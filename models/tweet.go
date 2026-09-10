package models

import (
	"time"
)

type Tweets struct {
	Id int `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
	Content string `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	LikeCount int `json:"like_count,omitempty" db:"-"`
	CommentCount int `json:"comment_count,omitempty" db:"-"`
}