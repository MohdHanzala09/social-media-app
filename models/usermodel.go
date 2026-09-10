package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)



type UserSchema struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	DOB time.Time `json:"dob"` 
}

type LoginSchema struct {
	Password string `json:"password"`
	Email string `json:"email"`
}

type Claims struct {
	Id int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}