package model

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("ZGFuc19tdWx0aV9wcm8")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
