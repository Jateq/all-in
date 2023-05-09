package main

import (
	"github.com/Jateq/all-in/controllers"
	"github.com/Jateq/all-in/middleware"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", controllers.Welcome)
	app.Post("/user/signup", controllers.SignUp)
	app.Post("/user/login", controllers.Login)
	app.Use(middleware.Authentication)
	app.Post("/user/addvault", controllers.AddVault)
	app.Post("/user/addfriend", controllers.AddFriend)
	app.Get("/user/friends", controllers.FriendList)
	app.Get("/user/vaults", controllers.Vaults)
	app.Post("/user/vault/:name", controllers.VaultToDos)
	//app.Get("/user/vault/day", controllers.DayToDos)
	log.Fatal(app.Listen(":4040"))

}
