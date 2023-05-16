package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Intialize Fiber
	app := fiber.New()

	// Routes
	movie := app.Group("/movie")

	movie.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("All movies")
	})

	movie.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Movie with id: " + c.Params("id"))
	})

	movie.Post("/", func(c *fiber.Ctx) error {
		return c.SendString("New movie")
	})

	movie.Put("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update movie with id: " + c.Params("id"))
	})

	movie.Delete("/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete movie with id: " + c.Params("id"))
	})

	// Start server
	app.Listen(":3000")
}
