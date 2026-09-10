package models

type LoginSchema struct {
	Password string `json:"password"`
	Email string `json:"email"`
}