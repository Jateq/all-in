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

var ToDoCollection *mongo.Collection = database.ToDoData(database.Client, "ToDoCollection")
var VaultCollection *mongo.Collection = database.VaultData(database.Client, "VaultCollection")
var UserCollection *mongo.Collection

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to ALL-IN")
}

func CreateToDo(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var toDo models.ToDo
	defer cancel()
	if err := c.BodyParser(&toDo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "error with JSON input"})
	}
	toDo.ToDoID = primitive.NewObjectID()
	toDo.Finished = time.Time{}
	_, anyErr := ToDoCollection.InsertOne(ctx, toDo)
	if anyErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "not inserted"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "Vault is created!"})
}

//func CreateDay(c *fiber.Ctx) error {
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
//	var day models.Day
//	defer cancel()
//	if err := c.BodyParser(&day); err != nil {
//		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "create a proper day plan"})
//	}
//	day.DayID = primitive.NewObjectID()
//	day.ToDos = make([]models.ToDo, 0)
//	day.EverythingDone = false
//	_, anyErr := DayCollection.InsertOne(ctx, day)
//	if anyErr != nil {
//		return c.Status(500).JSON(fiber.Map{"error": "not inserted to MongoDB"})
//	}
//	return c.Status(200).JSON(fiber.Map{"message": "Vault is created!"})
//
//}

func CreateVault(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var vault models.Vault
	if err := c.BodyParser(&vault); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "create a proper Vault"})
	}

	vault.VaultID = primitive.NewObjectID()
	vault.CreatedAt = time.Now()
	vault.StatusOverall = false
	vault.EachDay = make([]models.Day, 0)

	_, anyErr := VaultCollection.InsertOne(ctx, vault)
	if anyErr != nil {
		return c.Status(500).JSON(fiber.Map{"error": "not inserted to MongoDB"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Vault is created!"})
}

func Vaults(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var vaults []models.Vault
	cursor, err := VaultCollection.Find(ctx, bson.D{{}}) //again to return all the vaults
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "something went wrong, try later"})
	}
	err = cursor.All(ctx, &vaults)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "can't get vaults"})
	}
	defer cursor.Close(ctx)
	// If you don't close the cursor explicitly using cursor.Close(ctx),
	// it will remain open until the database connection is closed or garbage collected,
	// which can lead to resource leaks and connection pool exhaustion.
	if err = cursor.Err(); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "something went wrong"})
	}
	return c.Status(200).JSON(vaults)
}
