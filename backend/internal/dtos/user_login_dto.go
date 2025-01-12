package dtos

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaimsDto struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

type LoginRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	BaseResponseDto
	Token string      `json:"token,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}
