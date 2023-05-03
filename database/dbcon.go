package database

import (
	"context"
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
