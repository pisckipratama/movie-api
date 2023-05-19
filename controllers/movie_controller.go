package controllers

import (
	"strconv"

	model "github.com/asnur/movie-api/model"
	"github.com/gofiber/fiber/v2"
)

type MovieController struct {
	MovieModel model.MovieModel
}

func NewMovieController(MovieModel model.MovieModel) *MovieController {
	return &MovieController{
		MovieModel: MovieModel,
	}
}

// Get all movies
func (m *MovieController) All(c *fiber.Ctx) error {
	movies, err := m.MovieModel.All()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if len(movies) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "No movies found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "All movies",
		"data":    movies,
	})
}

// Get movie by id
func (m *MovieController) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	movie, err := m.MovieModel.Get(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Movie found",
		"data":    movie,
	})
}

// Create new movie
func (m *MovieController) Create(c *fiber.Ctx) error {
	var movies model.Movie

	if err := c.BodyParser(&movies); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	movie, err := m.MovieModel.Create(movies)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Movie created",
		"data":    movie,
	})
}

// Update movie
func (m *MovieController) Update(c *fiber.Ctx) error {
	var movies model.Movie

	if err := c.BodyParser(&movies); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Assign id from url param to model
	ID, _ := strconv.Atoi(c.Params("id"))
	movies.ID = uint(ID)

	movie, err := m.MovieModel.Update(movies)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Movie updated",
		"data":    movie,
	})
}

// Delete movie
func (m *MovieController) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	id, err := m.MovieModel.Delete(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Movie with ID " + strconv.Itoa(id) + " deleted",
	})

}
