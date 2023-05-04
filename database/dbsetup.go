package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func SetupDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect, try later")
		return nil
	}
	fmt.Println("Successfully connected to MongoDB")
	return client
}

var Client *mongo.Client = SetupDB()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var userCollection *mongo.Collection = client.Database("all-in").Collection(collectionName)
	return userCollection
}

func VaultData(client *mongo.Client, collectionName string) *mongo.Collection {
	var vaultCollection = client.Database("all-in").Collection(collectionName)
	return vaultCollection
}

func ToDoData(client *mongo.Client, collectionName string) *mongo.Collection {
	var toDoCollection *mongo.Collection = client.Database("all-in").Collection(collectionName)
	return toDoCollection
}
