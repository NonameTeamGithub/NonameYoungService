package jwtokens

import "github.com/dgrijalva/jwt-go"

type GenerateJWTParams struct {
	ExpiresIn int64
	UserId    string
}

type JWTClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}
