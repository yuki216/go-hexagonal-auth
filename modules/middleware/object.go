package middlewareJWT

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	bearer = "Bearer"
)

var (
	JWTSigningMethod = jwt.SigningMethodHS256
)

type JWTClaims struct {
	jwt.StandardClaims
	ID           int
	Email        string
	Name         string
}

type JWTRequest struct {
	ID           int
	Email        string
	Name         string
}
