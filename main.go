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
	app.Post("/user/createvault", controllers.CreateVault)
	app.Get("/user/vaults", controllers.Vaults)
	app.Post("/user/vault/addtodo", controllers.CreateToDo)
	//app.Get("/user/vault/day", controllers.DayToDos)
	log.Fatal(app.Listen(":4040"))

}
