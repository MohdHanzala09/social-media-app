package models

type Posts struct {
	Tweet string `json:"tweet"`
	UserID int `json:"userid"`
	UserEmail string `json:"email"`
}

type Likes struct {
	Likecount int `json:"likecount"`
	LikeByEmail string `json:"likebyemail"`
}

type Comments struct {
	Opinion string `json:"opinion"`
	CommentersEmail string `json:"commentersemail"`
}
