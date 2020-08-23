package repositories

import "github.com/inact25/movieplusplus/masters/apis/models"

type MoviesRepositories interface {
	GetAllMovie() ([]*models.MovieModels, error)
	GetMovieByID(id string) ([]*models.MovieModels, error)
	AddNewMovie(movies *models.MovieModels) (string, error)
}
