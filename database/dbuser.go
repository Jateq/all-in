package database

import (
	"context"
	"fmt"
	"github.com/Jateq/all-in/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func WorkWithUser(userCollection *mongo.Collection, userName string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	if err := userCollection.FindOne(ctx, bson.M{"username": userName}).Decode(&user); err != nil {
		fmt.Println(user)
		return user, err
	}

	return user, nil
}
