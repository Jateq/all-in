package controllers

import (
	"context"
	"github.com/Jateq/all-in/database"
	"github.com/Jateq/all-in/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var ToDoCollection *mongo.Collection = database.ToDoData(database.Client, "ToDoCollection")

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
	defer cancel()
	c.JSON("Successfully added")

	return nil
}

//func CreateDay(c *fiber.Ctx){
//}

