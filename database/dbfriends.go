package database

import (
	"context"
	"fmt"
	"github.com/Jateq/all-in/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// println's for my debugs

func AddFriendMongo(friend models.Friends, friendCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Insert the friend
	defer cancel()
	_, err := friendCollection.InsertOne(ctx, friend)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func sliceFriends(friendCollection *mongo.Collection, userID string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result []models.Friends

	searchFromDB, err := friendCollection.Find(ctx, bson.D{primitive.E{Key: "user_id", Value: userID}})
	if err != nil {
		fmt.Println("can't find friends")
		return nil, err
	}
	err = searchFromDB.All(ctx, &result)
	if err != nil {
		fmt.Println("can't assign to value")
		return nil, err
	}
	var friendSlice []string

	for _, tableFriend := range result {
		friendSlice = append(friendSlice, tableFriend.FriendID)
	}

	//fmt.Println(friendSlice)
	return friendSlice, nil

}

func FindFriends(userCollection, friendCollection *mongo.Collection, userID string) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var friends []models.User
	iDfriends, err := sliceFriends(friendCollection, userID)
	//fmt.Println("first init slice is", iDfriends)
	if err != nil {
		//fmt.Println("Error in slice func")
		return nil, err
	}
	for i := 0; i < len(iDfriends); i++ {
		var friend models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": iDfriends[i]}).Decode(&friend)
		if err != nil {
			//fmt.Println("can't find user with this id")
			return nil, err
		}

		if iDfriends[i] != userID {
			friends = append(friends, friend)
		}
	}
	//fmt.Println(friends)
	return friends, nil

}
