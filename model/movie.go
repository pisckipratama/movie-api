package model

type Movie struct {
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title" validate:"required"`
	Genre  string `json:"genre" validate:"required"`
	Year   int    `json:"year" validate:"required"`
	Poster string `json:"poster" validate:"required"`
}
