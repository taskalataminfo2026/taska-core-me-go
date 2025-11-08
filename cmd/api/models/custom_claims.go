package models

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
