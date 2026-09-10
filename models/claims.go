package models

import "github.com/golang-jwt/jwt/v5"


type Claims struct {
	Id int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}