package database

import (
	"context"
	"fmt"
	"github.com/Jateq/all-in/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ToDoCollection *mongo.Collection = ToDoData(Client, "ToDoCollection")

func DBInit() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var toDo models.ToDo
	toDo.ToDoName = "clean"
	toDo.Flag = false
	toDo.Finished = time.Now()

	ToDoCollection.InsertOne(ctx, toDo)
	defer cancel()
}

// just checking db connections

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
