package controllers

import (
	"context"
	"github.com/Jateq/all-in/database"
	"github.com/Jateq/all-in/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func AddFriend(c *fiber.Ctx) error {
	userID := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	var friend models.Friends
	if err := c.BodyParser(&friend); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": "Please make sure that you made a proper input"})
	}
	friend.StructID = primitive.NewObjectID()
	friend.UserID = userID
	if err := database.AddFriendMongo(friend, FriendCollection); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Can't add to database"})
	}
	ctx.Done()
	return c.Status(200).JSON(fiber.Map{
		"message": "successfully added friend",
	})

}

func FriendList(c *fiber.Ctx) error {
	userID := c.Query("id")
	friends, err := database.FindFriends(UserCollection, FriendCollection, userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "can't find your friends"})

	}
	return c.Status(200).JSON(friends)

}
