package usecases

import (
	"github.com/inact25/movieplusplus/masters/apis/models"
	"github.com/inact25/movieplusplus/masters/apis/repositories"
	"github.com/inact25/movieplusplus/utilities/validation"
)

type MovieUseCaseImpl struct {
	movieRepo repositories.MoviesRepositories
}

func (m MovieUseCaseImpl) GetAllMovie() ([]*models.MovieModels, error) {
	data, err := m.movieRepo.GetAllMovie()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m MovieUseCaseImpl) GetMovieByID(id string) ([]*models.MovieModels, error) {
	err := validation.CheckEmpty(id)
	if err != nil {
		return nil, err
	}
	menu, err := m.movieRepo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (m MovieUseCaseImpl) AddNewMovie(movies *models.MovieModels) (string, error) {
	err := validation.CheckEmpty(movies.MovieName, movies.MovieDuration, movies.MovieUrlPath, movies.MovieSynopsis)
	if err != nil {
		return "", err
	}
	movie, err := m.movieRepo.AddNewMovie(movies)
	if err != nil {
		return "", err
	}
	return movie, nil
}

func InitMovieUseCase(movies repositories.MoviesRepositories) MovieUseCases {
	return &MovieUseCaseImpl{movies}
}
