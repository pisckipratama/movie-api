package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title  string `json:"title" validate:"required" gorm:"column:title"`
	Genre  string `json:"genre" validate:"required" gorm:"column:genre"`
	Year   int    `json:"year" validate:"required" gorm:"column:year"`
	Poster string `json:"poster" validate:"required" gorm:"column:poster"`
}

func (m *Movie) TableName() string {
	return "movies"
}

type MovieModel struct {
	DB *gorm.DB
}

func NewMovieModel(db *gorm.DB) *MovieModel {
	return &MovieModel{
		DB: db,
	}
}

func (m *MovieModel) All() ([]Movie, error) {
	var movies []Movie
	if err := m.DB.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (m *MovieModel) Get(id int) (Movie, error) {
	var movie Movie
	if err := m.DB.First(&movie, id).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func (m *MovieModel) Create(movie Movie) (Movie, error) {
	if err := m.DB.Save(&movie).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func (m *MovieModel) Update(movie Movie) (Movie, error) {
	if err := m.DB.Save(&movie).Error; err != nil {
		return movie, err
	}
	return movie, nil
}

func (m *MovieModel) Delete(id int) (int, error) {
	if err := m.DB.Delete(&Movie{}, id).Error; err != nil {
		return 0, err
	}
	return id, nil
}
