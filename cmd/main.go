package main

import (
	"fmt"
	"log"

	"example.com/tomdoestech/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Hello from middleware")
		return c.Next()
	})

	app.Get("/healthcheck", healthcheck)

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	fmt.Println("Hello world")

	log.Fatal(app.Listen(":3000"))

}
