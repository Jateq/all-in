package controllers

import (
	"context"
	"github.com/Jateq/all-in/database"
	"github.com/Jateq/all-in/middleware"
	"github.com/Jateq/all-in/models"
	"github.com/Jateq/all-in/token"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var ToDoCollection *mongo.Collection = database.ToDoData(database.Client, "ToDoCollection")
var VaultCollection *mongo.Collection = database.VaultData(database.Client, "VaultCollection")
var UserCollection *mongo.Collection = database.UserData(database.Client, "UserCollection")
var FriendCollection = database.FriendsData(database.Client, "FriendsCollection")

var Validate = validator.New()

func Welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to ALL-IN")
}

func SignUp(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "create a proper user"})
	}
	validationErr := Validate.Struct(user)
	if validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": validationErr})
	}
	count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Something went wrong with server, try later"})
	}
	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User with this email already exists"})
	}
	hashedPassword := middleware.HashPassword(*user.Password)
	user.Password = &hashedPassword

	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	userToken, refreshToken, err := token.GenToken(*user.Email, *user.Username, user.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}
	user.Token = &userToken
	user.RefreshToken = &refreshToken
	user.Vaults = make([]models.Vault, 0)
	_, insertErr := UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong in registration"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Welcome to our community"})
}

func Login(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var user, foundUser models.User
	defer cancel()
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Write proper user struct"})
	}

	err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "can't find such email"})
	}
	passwordValid, msg := middleware.VerifyPassword(*user.Password, *foundUser.Password)
	if !passwordValid {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": msg})
	}
	userID := foundUser.UserID
	userToken, refreshToken, err := token.GenToken(*foundUser.Email, *foundUser.Username, foundUser.ID.Hex())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Can't generate token or refresh token"})
	}
	token.UpdateTokens(userToken, refreshToken, foundUser.ID.Hex())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello! Use variables below to get access to endpoints",
		"userID":  userID,
		"token":   userToken,
	})
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
