package main

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

type Dog struct {
	Name string
	Age  int
}

func main() {

	engine := html.New("templates", ".html")

	app := fiber.New(&fiber.Settings{Views: engine})

	app.Get("/", dogFunc)

	app.Listen(":8000")
}

func dogFunc(c *fiber.Ctx) {
	dog := &Dog{"boby", 3}
	_ = c.Render("index", fiber.Map{
		"Name": dog.Name, "Age": dog.Age,
	})
}
