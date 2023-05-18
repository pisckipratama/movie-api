package main

import (
	"github.com/asnur/movie-api/config"
	controller "github.com/asnur/movie-api/controllers"
	middleware "github.com/asnur/movie-api/middleware"
	"github.com/asnur/movie-api/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	// Intialize Fiber
	app := fiber.New()

	// Initialize controller
	movie_controller := new(controller.MovieController)
	user_controller := new(controller.UserController)

	// Routes User API
	user := app.Group("/user")
	user.Get("/login", middleware.ValidateField[model.User](), user_controller.Login)

	// Restricted routes
	app.Use(jwtware.New(config.JWTConfig))

	// Routes Movie API
	movie := app.Group("/movie")
	movie.Get("/", movie_controller.All)
	movie.Get("/:id", movie_controller.Get)
	movie.Post("/", middleware.ValidateField[model.Movie](), movie_controller.Create)
	movie.Put("/:id", middleware.ValidateField[model.Movie](), movie_controller.Update)
	movie.Delete("/:id", movie_controller.Delete)

	// Start server
	app.Listen(":3000")
}
