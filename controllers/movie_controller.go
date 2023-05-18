package controllers

import (
	"strconv"

	model "github.com/asnur/movie-api/model"
	"github.com/gofiber/fiber/v2"
)

type MovieController struct{}

// Get all movies
func (m *MovieController) All(c *fiber.Ctx) error {
	return c.SendString("All movies")
}

// Get movie by id
func (m *MovieController) Get(c *fiber.Ctx) error {
	return c.SendString("Movie with id: " + c.Params("id"))
}

// Create new movie
func (m *MovieController) Create(c *fiber.Ctx) error {
	var movie model.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	title := movie.Title
	genre := movie.Genre
	year := movie.Year
	poster := movie.Poster

	return c.SendString("value: " + title + " " + genre + " " + strconv.Itoa(year) + " " + poster)
}

// Update movie
func (m *MovieController) Update(c *fiber.Ctx) error {
	var movie model.Movie

	if err := c.BodyParser(&movie); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	title := movie.Title
	genre := movie.Genre
	year := movie.Year
	poster := movie.Poster

	return c.SendString("value: " + title + " " + genre + " " + strconv.Itoa(year) + " " + poster)
}

// Delete movie
func (m *MovieController) Delete(c *fiber.Ctx) error {
	return c.SendString("Delete movie with id: " + c.Params("id"))
}
