package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserId string
	Role   int8
	TokenNotify   string
	UserName   string
	jwt.StandardClaims
}
