package database

import (
	"context"
	"fmt"
	"github.com/Jateq/all-in/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func LinkVaultCommit(commitsCollection, userCollection *mongo.Collection, userID, dayID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

}

func LinkCommitTodos(todos []models.ToDo, commitsCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var commits models.Commits
	commits.DayID = primitive.NewObjectID()
	commits.DayNum = 0
	commits.ToDos = todos
	commits.EverythingDone = false
	_, err := commitsCollection.InsertOne(ctx, commits)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
