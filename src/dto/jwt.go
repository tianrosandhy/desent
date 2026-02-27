package dto

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
