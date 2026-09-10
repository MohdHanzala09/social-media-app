package models

import "time"

type Likes struct {
	Id int `json:"id" db:"id"`
	TweetId int `json:"tweet_id" db:"tweet_id"`
	UserId int `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}