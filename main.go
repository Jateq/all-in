package main

import (
	"github.com/Jateq/all-in/controllers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	app.Get("/", controllers.Welcome)
	app.Post("/user/vault/addtodo", controllers.CreateToDo)
	app.Get("user/vault/day", controllers.DayToDos)
	log.Fatal(app.Listen(":4040"))

}
