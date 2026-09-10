package models

import "time"

type UserSchema struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-" db:"password"`
	CaretedAt time.Time `json:"created_at"`
}
