package models

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JwtClaims struct {
	jwt.StandardClaims
	UserId    int   `jwt:"userId"`
	ExpiresAt int16 `jwt:"expiresAt"`
}
