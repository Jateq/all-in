package token

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

func GenToken(email, user, uid string) (signedToken, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email: email,
		User:  user,
		Uid:   uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(signedToken string) (*SignedDetails, error) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token is already expired")
	}
	return claims, nil
}

func UpdateTokens(signedToken, signedRefreshToken, userid string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	upsert := true
	filter := bson.M{"_id": userid}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{{
		Key: "$set", Value: updateObj},
	},
		&opt)

	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
}
