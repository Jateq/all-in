package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	Uid   string
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) { // not created yet.
	return claims, msg
}
