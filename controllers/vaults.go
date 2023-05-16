package controllers

import (
	"context"
	"github.com/Jateq/all-in/database"
	"github.com/Jateq/all-in/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func AddVault(c *fiber.Ctx) error {

	userID := c.Query("id")
	if userID == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "can't find this user"})
	}
	userVaultBID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("Internal server error")
	}

	var vault models.Vault
	vault.VaultID = primitive.NewObjectID()
	vault.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	vault.StatusOverall = false
	if err = c.BodyParser(&vault); err != nil {
		c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": "create a proper vault structure"})
	}

	var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	matchFilter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: userVaultBID}}}}
	unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$vaults"}}}}
	group := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value: &vault.VaultID}, {Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}
	pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't add"})
	}
	var vaultInfo []bson.M
	if err = pointcursor.All(ctx, &vaultInfo); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "can't get info about vaults"})
	}

	var size int32
	for _, num := range vaultInfo {
		count := num["count"]
		size = count.(int32)
	}
	if size < 3 {
		filter := bson.D{primitive.E{Key: "_id", Value: userVaultBID}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "vaults", Value: vault}}}}
		_, err := UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "can't update info about vaults"})
		}
	} else {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "Too many Vaults, please finish the rest"})
	}
	ctx.Done()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully created a Vault!",
		"Vault":   vault,
	})
}

func Vaults(c *fiber.Ctx) error {
	userID := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't find user's vault by token"})
	}

	vaults := user.Vaults

	return c.Status(200).JSON(vaults)
}

func VaultToDos(c *fiber.Ctx) error {
	userID := c.Query("id")
	vaultName := c.Params("name")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't find user's vault by token"})
	}

	var toDos []models.ToDo

	if err := c.BodyParser(&toDos); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with JSON input"})
	}

	for _, toDo := range toDos {
		toDo.ToDoID = primitive.NewObjectID()
		toDo.Finished = time.Time{}
		toDo.Flag = false
	}

	if err = database.LinkCommitTodos(toDos, CommitsCollection); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can't connect your todo plan"})
	}
	if err = database.LinkVaultCommit(CommitsCollection, UserCollection, userID, vaultName); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request",
			"more info": err})
	}
	return c.Status(200).JSON(fiber.Map{
		"message": "To do plan created,",
	})
}

func OneVault(c *fiber.Ctx) error {
	userID := c.Query("id")
	vaultName := c.Params("name")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't find user's vault by token"})
	}
	var reqVault models.Vault
	for _, reqVault = range user.Vaults {
		if *reqVault.VaultName == vaultName {
			return c.Status(200).JSON(reqVault)
		}
	}
	return c.Status(400).JSON(fiber.Map{"error": "can't find such vault"})

}

//func MyDay(c *fiber.Ctx) error {
//	userID := c.Query("id")
//	vaultName:
//	+c.Params("name")
//}

func VaultInfo(c *fiber.Ctx) error {
	userName := c.Params("username")
	var vaultName string = c.Params("vaultname")

	vault, err := database.FindVaultByVaultName(UserCollection, userName, vaultName)
	//fmt.Println(vault)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "something went wrong"})
	}

	return c.Status(200).JSON(vault)
}

//func FinishTodo(c *fiber.Ctx) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	c.Query()
//}
