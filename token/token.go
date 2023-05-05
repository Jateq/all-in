package token

import (
	"github.com/Jateq/all-in/database"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	User  string
	Uid   string
	jwt.StandardClaims
}

var UserData = database.UserData(database.Client, "UserCollection")
