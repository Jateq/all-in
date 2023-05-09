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

var commitID string

func LinkVaultCommit(commitsCollection, userCollection *mongo.Collection, userID, vaultName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return err
	}

	vaults := user.Vaults
	for i := range vaults {
		if vaultName == *vaults[i].VaultName {
			vaults[i].DayPlan = commitID
			_, updateErr := userCollection.UpdateOne(ctx, bson.M{"user_id": userID}, bson.M{"$set": bson.M{"vaults": vaults}})
			if updateErr != nil {
				return updateErr
			}
			return nil
		}
		fmt.Println("no such vault")
		return err
	}

	return nil
}

func LinkCommitTodos(todos []models.ToDo, commitsCollection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var commits models.Commits
	commits.DayID = primitive.NewObjectID()
	commits.CommitID = commits.DayID.Hex()
	commitID = commits.CommitID
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
